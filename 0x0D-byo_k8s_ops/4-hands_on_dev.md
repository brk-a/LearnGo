# hands-on development

<div style="text-align: justify;">

## 0. Prerequisites (2 minutes)

```bash
    # 1. AWS credentials (IAM user with EC2 permissions)
    aws sts get-caller-identity  # Verify access

    # 2. Go dependencies
    go mod tidy

    # 3. Verify
    kubectl get mysqlcluster  # Should show alx-nairobi-prod running
```

## 1. Implementing AWS SDK Integration

### 1.1. AWS Client Factory

* Create `internal/aws/client.go`:

```go
    package aws

    import (
        "context"
        "fmt"
        
        "github.com/aws/aws-sdk-go-v2/aws"
        "github.com/aws/aws-sdk-go-v2/config"
        "github.com/aws/aws-sdk-go-v2/service/ec2"
        "github.com/aws/aws-sdk-go-v2/service/ec2/types"
        "sigs.k8s.io/controller-runtime/pkg/log"
    )

    // EC2Client manages AWS EC2 operations for the operator
    type EC2Client struct {
        client *ec2.Client
        log    log.Logger
    }

    // NewEC2Client creates AWS client from environment/default credentials
    func NewEC2Client(ctx context.Context, region string) (*EC2Client, error) {
        cfg, err := config.LoadDefaultConfig(ctx, 
            config.WithRegion(region),
            config.WithLogger(awsLogger{}),
        )
        if err != nil {
            return nil, fmt.Errorf("failed to load AWS config: %w", err)
        }
        
        return &EC2Client{
            client: ec2.NewFromConfig(cfg),
            log:    log.FromContext(ctx),
        }, nil
    }

    // ProvisionMySQLE instances creates EC2 instances for MySQL cluster
    func (c *EC2Client) ProvisionMySQLInstances(ctx context.Context, clusterName string, replicas int32, config Config) ([]string, error) {
        c.log.Info("Provisioning EC2 instances", "cluster", clusterName, "replicas", replicas, "type", config.InstanceType)
        
        input := &ec2.RunInstancesInput{
            ImageId:      aws.String(config.AMI),
            InstanceType: config.InstanceType,
            MinCount:     aws.Int32(replicas),
            MaxCount:     aws.Int32(replicas),
            SubnetId:     aws.String(config.SubnetID),
            SecurityGroupIds: []string{config.SecurityGroupID},
            TagSpecifications: []types.TagSpecification{{
                ResourceType: types.ResourceTypeInstance,
                Tags: []types.Tag{
                    {
                        Key:   aws.String("Name"),
                        Value: aws.String(fmt.Sprintf("mysql-%s", clusterName)),
                    },
                    {
                        Key:   aws.String("kubernetes.io/cluster/nairobi-mysql"),
                        Value: aws.String("owned"),
                    },
                    {
                        Key:   aws.String("app.kubernetes.io/managed-by"),
                        Value: aws.String("nairobi-mysql-operator"),
                    },
                },
            }},
        }
        
        result, err := c.client.RunInstances(ctx, input)
        if err != nil {
            return nil, fmt.Errorf("failed to run instances: %w", err)
        }
        
        instanceIDs := make([]string, 0, len(result.Instances))
        for _, inst := range result.Instances {
            instanceIDs = append(instanceIDs, *inst.InstanceId)
            c.log.Info("Launched EC2 instance", "id", *inst.InstanceId, "state", *inst.State.Name)
        }
        
        return instanceIDs, nil
    }

    // WaitForRunning polls until all instances are running
    func (c *EC2Client) WaitForRunning(ctx context.Context, instanceIDs []string) error {
        c.log.Info("Waiting for instances to become running", "count", len(instanceIDs))
        
        input := &ec2.DescribeInstancesInput{
            InstanceIds: instanceIDs,
        }
        
        waiter := ec2.NewWaitForInstanceRunningWaiter(c.client)
        
        return waiter.Wait(ctx, input, waiter.MaxAttempts(60), waiter.Delay(10*time.Second))
    }

    // TerminateInstances destroys EC2 instances during cleanup
    func (c *EC2Client) TerminateInstances(ctx context.Context, instanceIDs []string) error {
        c.log.Info("Terminating EC2 instances", "count", len(instanceIDs))
        
        input := &ec2.TerminateInstancesInput{
            InstanceIds: instanceIDs,
        }
        
        _, err := c.client.TerminateInstances(ctx, input)
        if err != nil {
            return fmt.Errorf("failed to terminate instances: %w", err)
        }
        
        c.log.Info("Termination initiated successfully")
        return nil
    }

    // Config holds AWS infrastructure configuration
    type Config struct {
        AMI              string
        InstanceType     types.InstanceType
        SubnetID         string
        SecurityGroupID  string
    }
```

### 1.2. Enhanced CRD Specification

* Update `api/v1/mysqlcluster_types.go` with AWS fields:

```go
    // Add to MySQLClusterSpec
    type AWSConfig struct {
        AMI              string   `json:"ami"`
        InstanceType     string   `json:"instanceType"`
        SubnetID         string   `json:"subnetId"`
        SecurityGroupIDs []string `json:"securityGroupIds"`
        KeyName          string   `json:"keyName,omitempty"`
    }

    type MySQLClusterSpec struct {
        // ... existing fields ...
        AWS AWSConfig `json:"aws,omitempty"`
    }
```

* **Note**: **Never hardcode credentials** - operator uses IAM roles/EC2 metadata service in production.

## 2. finalisers for Graceful Cleanup

* **Note**: finalisers ensure **EC2 instances are terminated** when CR is deleted.

* Update controller `Reconcile()`:

```go
    func (r *MySQLClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) ctrl.Result {
        // ... existing code ...
        
        // Handle deletion finaliser FIRST
        if cluster.DeletionTimestamp != nil {
            return r.handleDeletion(ctx, cluster)
        }
        
        // Normal reconciliation...
        return r.reconcileNormal(ctx, cluster)
    }

    func (r *MySQLClusterReconciler) handleDeletion(ctx context.Context, cluster *mysqlv1.MySQLCluster) ctrl.Result {
        log := log.FromContext(ctx)
        
        // Check if finaliser present
        if !controllerutil.Containsfinaliser(cluster, mysqlfinaliser) {
            return ctrl.Result{}
        }
        
        // Cleanup AWS resources
        if len(cluster.Status.AWSInstanceIDs) > 0 {
            awsClient, err := aws.NewEC2Client(ctx, "us-east-1")
            if err != nil {
                log.Error(err, "Failed to create AWS client during cleanup")
                return ctrl.Result{RequeueAfter: 5 * time.Minute}
            }
            
            if err := awsClient.TerminateInstances(ctx, cluster.Status.AWSInstanceIDs); err != nil {
                log.Error(err, "Failed to terminate EC2 instances")
                return ctrl.Result{RequeueAfter: 2 * time.Minute}
            }
            
            log.Info("EC2 instances terminated successfully")
            cluster.Status.AWSInstanceIDs = nil
            cluster.Status.Phase = "Terminating"
        }
        
        // Remove finaliser when clean
        controllerutil.Removefinaliser(cluster, mysqlfinaliser)
        if err := r.Status().Update(ctx, cluster); err != nil {
            return ctrl.Result{}, err
        }
        
        log.Info("Cleanup complete, finaliser removed")
        return ctrl.Result{}
    }
```

## 3. Complete Controller with AWS Integration

* **Full `controllers/mysqlcluster_controller.go`**:

```go
    // Add to top of Reconcile()
    const mysqlfinaliser = "mysql.alx.ke/finaliser"

    func (r *MySQLClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) ctrl.Result {
        log := log.FromContext(ctx)
        cluster := &mysqlv1.MySQLCluster{}
        
        if err := r.Get(ctx, req.NamespacedName, cluster); err != nil {
            return ctrl.Result{}, client.IgnoreNotFound(err)
        }
        
        log.Info("Reconciling", "cluster", cluster.Name, "generation", cluster.Generation)
        
        // Handle deletion
        if cluster.DeletionTimestamp != nil {
            return r.handleDeletion(ctx, cluster)
        }
        
        // Add finaliser on first reconciliation
        if !controllerutil.Containsfinaliser(cluster, mysqlfinaliser) {
            controllerutil.Addfinaliser(cluster, mysqlfinaliser)
            if err := r.Update(ctx, cluster); err != nil {
                return ctrl.Result{}, err
            }
            log.Info("Added finaliser")
            return ctrl.Result{RequeueAfter: 10 * time.Second}
        }
        
        // Provision AWS infrastructure
        if cluster.Status.Phase == "" || cluster.Status.Phase == "Pending" {
            if err := r.provisionAWSInfrastructure(ctx, cluster); err != nil {
                log.Error(err, "AWS provisioning failed")
                cluster.Status.Phase = "ProvisioningFailed"
                cluster.Status.Conditions = append(cluster.Status.Conditions, 
                    mysqlv1.Condition{
                        Type:    "Provisioned",
                        Status:  "False",
                        Reason:  "AWSAPIError",
                        Message: err.Error(),
                    })
                r.Status().Update(ctx, cluster)
                return ctrl.Result{RequeueAfter: 30 * time.Second}, err
            }
        }
        
        // Normal K8s resources...
        r.reconcileKubernetesResources(ctx, cluster)
        
        return ctrl.Result{RequeueAfter: 30 * time.Second}
    }

    func (r *MySQLClusterReconciler) provisionAWSInfrastructure(ctx context.Context, cluster *mysqlv1.MySQLCluster) error {
        awsClient, err := aws.NewEC2Client(ctx, "us-east-1")
        if err != nil {
            return err
        }
        
        config := aws.Config{
            AMI:              cluster.Spec.AWS.AMI,
            InstanceType:     types.InstanceTypeT3Medium,
            SubnetID:         cluster.Spec.AWS.SubnetID,
            SecurityGroupID:  cluster.Spec.AWS.SecurityGroupIDs[0],
        }
        
        instanceIDs, err := awsClient.ProvisionMySQLInstances(ctx, cluster.Name, cluster.Spec.Replicas, config)
        if err != nil {
            return err
        }
        
        // Wait for instances to be ready
        if err := awsClient.WaitForRunning(ctx, instanceIDs); err != nil {
            return fmt.Errorf("instances failed to start: %w", err)
        }
        
        // Update status
        cluster.Status.AWSInstanceIDs = instanceIDs
        cluster.Status.Phase = "Provisioned"
        
        log.FromContext(ctx).Info("AWS infrastructure provisioned", "instanceCount", len(instanceIDs))
        return r.Status().Update(ctx, cluster)
    }
```

## 4. Test Live AWS Provisioning

* **Create test CR**:

```bash
    cat > aws-mysql-test.yaml << 'EOF'
    apiVersion: mysql.alx.ke/v1
    kind: MySQLCluster
    metadata:
    name: aws-test-cluster
    spec:
    replicas: 2
    aws:
        ami: "ami-0c02fb55956c7d316"  # Ubuntu 22.04 us-east-1
        instanceType: "t3.micro"       # Free tier!
        subnetId: "subnet-12345678"    # Your VPC subnet
        securityGroupIds: ["sg-12345678"]
    geoAffinity: "nairobi"
    EOF

    kubectl apply -f aws-mysql-test.yaml
```

* **Monitor provisioning**:

```bash
    # Watch CR status
    kubectl get mysqlcluster aws-test-cluster -w

    # Controller logs
    kubectl logs -n nairobi-mysql-system deployment/nairobi-mysql-controller-manager-manager -f

    # AWS CLI verification
    aws ec2 describe-instances --filters "Name=tag:Name,Values=mysql-aws-test-cluster"
```

* **Expected flow**:

```plaintext
    09:15 - kubectl apply → Status: Pending
    09:16 - "Provisioning EC2 instances" (controller log)
    09:17 - "Launched EC2 instance i-0123456789abcdef0" 
    09:18 - "Waiting for instances to become running"
    09:20 - Status: Provisioned, AWSInstanceIDs: [i-xxx, i-yyy]
```

## 5. Test Cleanup finaliser

```bash
    kubectl delete mysqlcluster aws-test-cluster

    # Watch finaliser
    kubectl get mysqlcluster aws-test-cluster -w
    # finaliser blocks deletion...

    # Controller logs show:
    # "Terminating EC2 instances"
    # "Cleanup complete, finaliser removed"

    # CR disappears + EC2 instances terminated
    aws ec2 describe-instances --filters "Name=instance-state-name,Values=terminated" --query "Reservations[*].Instances[*].InstanceId"
```

***

## 6. version control and clean-up

```bash
    git add .
    git commit -m "feat: day3 AWS EC2 integration + finalisers

    - Real EC2 provisioning (RunInstances → WaitUntilRunning)
    - finaliser cleanup (TerminateInstances on CR deletion)
    - AWS SDK v2 integration with structured logging
    - Production error handling + exponential backoff
    - Status reporting with AWSInstanceIDs array"
    git push
```


* **Clean up test instances**:

```bash
    kubectl delete mysqlcluster aws-test-cluster --force --grace-period=0
```

</div>
