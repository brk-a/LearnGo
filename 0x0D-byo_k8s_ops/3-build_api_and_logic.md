# build API and logic

<div style="text-align: justify;">

## 0. Creating & Customising Your First API

* **Note**: Kubebuilder has already scaffolded the basic API structure from [2-set_up_env][def]; now we extend it with **Nairobi-specific cloud optimisations** for EC2 instance sizing and geo-aware deployments.

### 0.1. EC2 Instance Types & Nairobi Specifications

* Navigate to `api/v1/mysqlcluster_types.go` and replace the default types with production-ready specifications:

```go
    package v1

    import (
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    )

    // Nairobi sizing constants - optimised for cost/performance
    const (
        InstanceClassMicro  = "t3.micro"   // Development workshops
        InstanceClassMedium = "t3.medium"  // Staging environments
        InstanceClassLarge  = "m5.large"   // Production workloads
        InstanceClassXL     = "r6g.xlarge" // Analytics heavy
    )

    // +kubebuilder:object:root=true
    // +kubebuilder:subresource:status
    type MySQLCluster struct {
        metav1.TypeMeta   `json:",inline"`
        metav1.ObjectMeta `json:"metadata,omitempty"`

        Spec   MySQLClusterSpec   `json:"spec,omitempty"`
        Status MySQLClusterStatus `json:"status,omitempty"`
    }

    // +kubebuilder:object:root=true
    type MySQLClusterList struct {
        metav1.TypeMeta `json:",inline"`
        metav1.ListMeta `json:"metadata,omitempty"`
        Items           []MySQLCluster `json:"items"`
    }

    // MySQLClusterSpec defines the **desired state** - what users declare
    type MySQLClusterSpec struct {
        // +kubebuilder:default=1
        Replicas int32 `json:"replicas,omitempty"`
        
        // Cloud provider configuration with auto-sizing
        CloudConfig CloudConfig `json:"cloudConfig"`
        
        // Persistent storage specification
        Storage MySQLStorage `json:"storage"`
        
        // Automated backup configuration
        Backup BackupConfig `json:"backup,omitempty"`
        
        // Kenyan timezone/geo preference
        GeoAffinity string `json:"geoAffinity,omitempty"`
    }

    // CloudConfig encapsulates AWS EC2 optimisations
    type CloudConfig struct {
        Provider     string `json:"provider"`     // "aws", "gcp", "azure"
        InstanceClass string `json:"instanceClass,omitempty"`
        CPURequest   string `json:"cpuRequest,omitempty"`
        MemoryRequest string `json:"memoryRequest,omitempty"`
    }

    // MySQLStorage follows CSI storage best practices
    type MySQLStorage struct {
        Size         string `json:"size"`         // "20Gi", "100Gi"
        StorageClass string `json:"storageClass"` // "nairobi-local-fast"
        IOPS         *int64 `json:"iops,omitempty"`
    }

    // BackupConfig integrates with S3/Velero
    type BackupConfig struct {
        Enabled  bool   `json:"enabled,omitempty"`
        Schedule string `json:"schedule,omitempty"` // "0 2 * * *" (3AM EAT)
        S3Bucket string `json:"s3Bucket,omitempty"`
    }

    // MySQLClusterStatus reflects the **observed state** - controller reports back
    type MySQLClusterStatus struct {
        Phase string `json:"phase,omitempty"` // Pending, Running, Failed
        
        Replicas    int32 `json:"replicas,omitempty"`
        ReplicasReady int32 `json:"replicasReady,omitempty"`
        
        Endpoints Endpoints `json:"endpoints,omitempty"`
        
        Conditions []metav1.Condition `json:"conditions,omitempty"`
    }

    type Endpoints struct {
        ReadEndpoint  string `json:"readEndpoint,omitempty"`
        WriteEndpoint string `json:"writeEndpoint,omitempty"`
    }

    func init() {
        SchemeBuilder.Register(&MySQLCluster{}, &MySQLClusterList{})
    }
```

* **Note**: The `+kubebuilder` comments are **marker directives** that `controller-gen` uses to auto-generate CRDs, RBAC, and client code. `Spec` represents **desired state** (user input); `Status` shows **observed state** (controller output).

* **Regenerate manifests**:

```bash
    make generate manifests
    kubectl apply -f config/crd/bases/mysql.alx.ke_mysqlclusters.yaml
```

### 0.2. TypeMeta and ObjectMeta Explained

* **TypeMeta** defines the **API contract** (immutable):

```go
    type TypeMeta struct {
        APIVersion string `json:"apiVersion"` // "mysql.alx.ke/v1"
        Kind       string `json:"kind"`       // "MySQLCluster"
    }
```

* **ObjectMeta** provides **Kubernetes metadata** (mutable):

```go
    type ObjectMeta struct {
        Name            string            `json:"name"`              // "alx-nairobi-prod"
        Namespace       string            `json:"namespace"`         // "default"
        Labels          map[string]string `json:"labels,omitempty"`  // app.kubernetes.io/name
        Finalizers      []string          `json:"finalizers,omitempty"` // Cleanup hooks
        ResourceVersion string            `json:"resourceVersion"`   // etcd optimistic lock
        Generation      int64             `json:"generation"`        // Spec hash changes
    }
```

* **Note**: `Generation` increments when `.spec` changes, triggering reconciliation even if `.status` matches.

### 0.3. Internal Controller Logic Breakdown

* Replace `controllers/mysqlcluster_controller.go` with production logic:

```go
    package controllers

    import (
        "context"
        "fmt"
        "time"
        
        appsv1 "k8s.io/api/apps/v1"
        corev1 "k8s.io/api/core/v1"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/apimachinery/pkg/runtime"
        ctrl "sigs.k8s.io/controller-runtime"
        "sigs.k8s.io/controller-runtime/pkg/client"
        "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
        "sigs.k8s.io/controller-runtime/pkg/log"
        
        mysqlv1 "github.com/yourusername/nairobi-mysql-operator/api/v1"
    )

    type MySQLClusterReconciler struct {
        client.Client
        Scheme *runtime.Scheme
    }

    // +kubebuilder:rbac:groups=mysql.alx.ke,resources=mysqlclusters,verbs=get;list;watch;create;update;patch;delete
    // +kubebuilder:rbac:groups=mysql.alx.ke,resources=mysqlclusters/status,verbs=get;update;patch
    // +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
    // +kubebuilder:rbac:groups=core,resources=services;persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete

    func (r *MySQLClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) ctrl.Result {
        log := log.FromContext(ctx)
        
        // Phase 1: Fetch the MySQLCluster instance
        cluster := &mysqlv1.MySQLCluster{}
        if err := r.Get(ctx, req.NamespacedName, cluster); err != nil {
            return ctrl.Result{}, client.IgnoreNotFound(err)
        }
        
        log.Info("Reconciling MySQLCluster", 
            "name", cluster.Name, 
            "namespace", cluster.Namespace,
            "generation", cluster.GetGeneration())
        
        // Phase 2: Apply intelligent defaults (Nairobi optimisations)
        if err := r.defaultNairobiConfig(cluster); err != nil {
            log.Error(err, "Failed to apply defaults")
            return ctrl.Result{RequeueAfter: 30 * time.Second}, err
        }
        
        // Phase 3: Reconcile child resources (idempotent)
        if err := r.reconcileStatefulSet(ctx, cluster); err != nil {
            log.Error(err, "StatefulSet reconciliation failed")
            return ctrl.Result{RequeueAfter: 10 * time.Second}, err
        }
        
        if err := r.reconcileHeadlessService(ctx, cluster); err != nil {
            return ctrl.Result{RequeueAfter: 10 * time.Second}, err
        }
        
        if cluster.Spec.Backup.Enabled {
            if err := r.reconcileBackupCronJob(ctx, cluster); err != nil {
                log.Error(err, "Backup job failed")
                return ctrl.Result{RequeueAfter: 30 * time.Second}, err
            }
        }
        
        // Phase 4: Update observed status
        if err := r.updateClusterStatus(ctx, cluster); err != nil {
            log.Error(err, "Status update failed")
            return ctrl.Result{}, err
        }
        
        log.Info("MySQLCluster successfully reconciled", "name", cluster.Name)
        return ctrl.Result{}  // No immediate requeue needed
    }

    // defaultNairobiConfig applies geo-aware defaults for Kenyan deployments
    func (r *MySQLClusterReconciler) defaultNairobiConfig(cluster *mysqlv1.MySQLCluster) error {
        // Auto-size instance class based on replica count
        if cluster.Spec.CloudConfig.InstanceClass == "" {
            switch cluster.Spec.Replicas {
            case 1:
                cluster.Spec.CloudConfig.InstanceClass = "t3.micro"   // Workshop
            case 3:
                cluster.Spec.CloudConfig.InstanceClass = "t3.medium"  // Staging
            default:
                cluster.Spec.CloudConfig.InstanceClass = "m5.large"   // Production
            }
        }
        
        // Default to Nairobi timezone/geo
        if cluster.Spec.GeoAffinity == "" {
            cluster.Spec.GeoAffinity = "nairobi"
        }
        
        // Ensure AWS provider for EC2 optimisations
        if cluster.Spec.CloudConfig.Provider == "" {
            cluster.Spec.CloudConfig.Provider = "aws"
        }
        
        return nil
    }

    // reconcileStatefulSet creates/updates the MySQL StatefulSet idempotently
    func (r *MySQLClusterReconciler) reconcileStatefulSet(ctx context.Context, cluster *mysqlv1.MySQLCluster) error {
        sts := r.generateStatefulSet(cluster)
        
        // Create or patch existing StatefulSet
        err := r.reconcileResource(ctx, sts, cluster)
        if err != nil {
            return fmt.Errorf("statefulset reconciliation failed: %w", err)
        }
        
        // Update cluster status with StatefulSet readiness
        cluster.Status.Replicas = *sts.Spec.Replicas
        if sts.Status.ReadyReplicas > 0 {
            cluster.Status.ReplicasReady = sts.Status.ReadyReplicas
            cluster.Status.Phase = "Running"
        }
        
        return nil
    }

    // generateStatefulSet creates the StatefulSet manifest from cluster spec
    func (r *MySQLClusterReconciler) generateStatefulSet(cluster *mysqlv1.MySQLCluster) *appsv1.StatefulSet {
        labels := map[string]string{
            "app":                           cluster.Name,
            "app.kubernetes.io/name":        "mysql",
            "app.kubernetes.io/managed-by":  "nairobi-mysql-operator",
            "app.kubernetes.io/instance":    cluster.Name,
        }
        
        return &appsv1.StatefulSet{
            ObjectMeta: metav1.ObjectMeta{
                Name:      cluster.Name,
                Namespace: cluster.Namespace,
                Labels:    labels,
            },
            Spec: appsv1.StatefulSetSpec{
                ServiceName: fmt.Sprintf("%s-headless", cluster.Name),
                Replicas:    &cluster.Spec.Replicas,
                Selector:    &metav1.LabelSelector{MatchLabels: labels},
                Template: corev1.PodTemplateSpec{
                    ObjectMeta: metav1.ObjectMeta{Labels: labels},
                    Spec: corev1.PodSpec{
                        Containers: []corev1.Container{{
                            Name:  "mysql",
                            Image: "mysql:8.0.35",
                            Ports: []corev1.ContainerPort{{ContainerPort: 3306, Name: "mysql"}},
                            Env: []corev1.EnvVar{
                                {
                                    Name:  "MYSQL_ROOT_PASSWORD",
                                    Value: "alx-nairobi-2026!",
                                },
                            },
                            Resources: r.calculateCloudResources(cluster.Spec.CloudConfig),
                            VolumeMounts: []corev1.VolumeMount{{
                                Name:      "mysql-data",
                                MountPath: "/var/lib/mysql",
                            }},
                            ReadinessProbe: &corev1.Probe{
                                ProbeHandler: corev1.ProbeHandler{
                                    TCPSocket: &corev1.TCPSocketAction{Port: 3306},
                                },
                                InitialDelaySeconds: 30,
                                PeriodSeconds:       10,
                            },
                        }},
                        Volumes: []corev1.Volume{{
                            Name: "mysql-data",
                            VolumeSource: corev1.VolumeSource{
                                EmptyDir: &corev1.EmptyDirVolumeSource{},
                            },
                        }},
                    },
                },
            },
        }
    }

    // calculateCloudResources maps EC2 instance types to Kubernetes resources
    func (r *MySQLClusterReconciler) calculateCloudResources(config mysqlv1.CloudConfig) corev1.ResourceRequirements {
        resourceMap := map[string]corev1.ResourceRequirements{
            "t3.micro": {
                Requests: corev1.ResourceList{
                    "cpu":    resource.MustParse("250m"),
                    "memory": resource.MustParse("512Mi"),
                },
                Limits: corev1.ResourceList{
                    "cpu":    resource.MustParse("500m"),
                    "memory": resource.MustParse("1Gi"),
                },
            },
            "t3.medium": {
                Requests: corev1.ResourceList{
                    "cpu":    resource.MustParse("1"),
                    "memory": resource.MustParse("2Gi"),
                },
                Limits: corev1.ResourceList{
                    "cpu":    resource.MustParse("2"),
                    "memory": resource.MustParse("4Gi"),
                },
            },
            "m5.large": {
                Requests: corev1.ResourceList{
                    "cpu":    resource.MustParse("2"),
                    "memory": resource.MustParse("8Gi"),
                },
            },
        }
        
        if res, exists := resourceMap[config.InstanceClass]; exists {
            return res
        }
        
        // Default production sizing
        return resourceMap["t3.medium"]
    }

    // reconcileResource is the idempotent create/update helper
    func (r *MySQLClusterReconciler) reconcileResource(ctx context.Context, obj client.Object, owner client.Object) error {
        log := log.FromContext(ctx)
        
        err := controllerutil.SetControllerReference(owner, obj, r.Scheme)
        if err != nil {
            return err
        }
        
        var existing client.Object
        switch t := obj.(type) {
        case *appsv1.StatefulSet:
            existing = &appsv1.StatefulSet{}
        case *corev1.Service:
            existing = &corev1.Service{}
        }
        
        key := client.ObjectKeyFromObject(obj)
        err = r.Get(ctx, key, existing)
        
        if err == nil {
            // UPDATE: Strategic Merge Patch (idempotent)
            patch := client.MergeFrom(existing)
            err = r.Patch(ctx, obj, patch)
            log.Info("Patched resource", "kind", obj.GetObjectKind().GroupVersionKind().Kind, "name", obj.GetName())
        } else if client.IgnoreNotFound(err) == nil {
            // CREATE
            err = r.Create(ctx, obj)
            log.Info("Created resource", "kind", obj.GetObjectKind().GroupVersionKind().Kind, "name", obj.GetName())
        }
        
        return err
    }

    // SetupWithManager configures the controller on the manager
    func (r *MySQLClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
        return ctrl.NewControllerManagedBy(mgr).
            For(&mysqlv1.MySQLCluster{}).
            Owns(&appsv1.StatefulSet{}).
            Owns(&corev1.Service{}).
            Complete(r)
    }
```

* **Note**: The controller follows the **pure function pattern**: `reconcile(MySQLCluster.Spec) → Kubernetes Resources`. Repeated calls converge to identical state.

## 1. Production Manager Architecture

* **Note**: The `controller-runtime` Manager orchestrates the entire operator lifecycle:

```plaintext
    Manager Responsibilities:
    ├── REST Client (shared informer cache)
    ├── Scheme (type registration)  
    ├── Leader Election (HA)
    ├── HTTPS Server (port 9443 - webhooks)
    ├── Metrics Server (port 8080 - Prometheus)
    ├── Health Checks (/healthz, /readyz, /livez)
    └── Graceful shutdown
```

* **Enhanced main.go** with full production features:

```go
    package main

    import (
        "context"
        "flag"
        "os"
        
        "k8s.io/apimachinery/pkg/runtime"
        ctrl "sigs.k8s.io/controller-runtime"
        "sigs.k8s.io/controller-runtime/pkg/healthz"
        "sigs.k8s.io/controller-runtime/pkg/log/zap"
        
        mysqlv1 "github.com/yourusername/nairobi-mysql-operator/api/v1"
        "github.com/yourusername/nairobi-mysql-operator/controllers"
    )

    var (
        scheme   = runtime.NewScheme()
        setupLog = ctrl.Log.WithName("setup")
    )

    func init() {
        mysqlv1.AddToScheme(scheme)
        // Note: controller-tools will auto-add other schemes
    }

    func main() {
        var (
            metricsAddr           = flag.String("metrics-addr", ":8080", "The address the metric endpoint binds to")
            enableLeaderElection  = flag.Bool("leader-elect", false, "Enable leader election")
            probeAddr             = flag.String("health-probe-bind-addr", ":8081", "Health probe bind address")
        )
        flag.Parse()
        
        // Structured logging with zap
        opts := zap.Options{Development: true}
        ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
        
        // Create production manager
        mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
            Scheme:                 scheme,
            MetricsBindAddress:     *metricsAddr,
            Port:                   9443,        // Webhook server
            HealthProbeBindAddress: *probeAddr,  // /healthz /readyz /livez
            LeaderElection:         *enableLeaderElection,
            LeaderElectionID:       "nairobi-mysql.alx.ke",
            // CertDir:                "/tmp/k8s-webhook-server-serving-certs", // Auto-managed
        })
        if err != nil {
            setupLog.Error(err, "unable to start manager")
            os.Exit(1)
        }
        
        // Register health checks
        if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
            panic(err)
        }
        if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
            panic(err)
        }
        
        // Register controllers
        if err = (&controllers.MySQLClusterReconciler{
            Client: mgr.GetClient(),
            Scheme: mgr.GetScheme(),
        }).SetupWithManager(mgr); err != nil {
            setupLog.Error(err, "unable to create controller")
            os.Exit(1)
        }
        
        // Note: Leader election creates Lease CR for HA failover
        setupLog.Info("starting manager", 
            "leaderElection", *enableLeaderElection, 
            "metricsAddr", *metricsAddr,
            "probeAddr", *probeAddr)
        
        // Graceful shutdown
        if err := mgr.Start(context.TODO()); err != nil {
            setupLog.Error(err, "problem running manager")
            os.Exit(1)
        }
    }
```

## 2. Testing Your Production Operator

* **Note**: Test the full reconciliation loop:

```bash
    # 1. Rebuild everything
    make generate manifests fmt vet

    # 2. Restart controller (in first terminal)
    make run

    # 3. Deploy production workload (second terminal)
    cat > nairobi-prod.yaml << 'EOF'
    apiVersion: mysql.alx.ke/v1
    kind: MySQLCluster
    metadata:
    name: alx-nairobi-prod
    namespace: default
    spec:
    replicas: 3
    cloudConfig:
        provider: aws
    storage:
        size: 50Gi
        storageClass: nairobi-local-fast
    backup:
        enabled: true
        schedule: "0 2 * * *"
    geoAffinity: "nairobi"
    EOF

    kubectl apply -f nairobi-prod.yaml
```

* **Monitor reconciliation** (third terminal):

```bash
    # Watch CR status
    kubectl get mysqlcluster alx-nairobi-prod -w

    # Check child resources
    kubectl get statefulset,service,pvc -l app=alx-nairobi-prod

    # Controller logs
    kubectl logs deployment/nairobi-mysql-controller-manager-manager -n nairobi-mysql-system -f | grep alx-nairobi-prod

    # Metrics endpoint
    curl http://localhost:8080/metrics | grep mysqlcluster
```

* **Expected output**:

```plaintext
    $ kubectl get mysqlcluster alx-nairobi-prod
    NAME                REPLICAS   READY   STATUS   AGE
    alx-nairobi-prod    3          3/3     Running  2m40s

    $ kubectl get statefulset alx-nairobi-prod
    NAME                READY   AGE
    alx-nairobi-prod    3/3     2m40s
```

* **Note**: **Success criteria**:
- ✅ CRDs updated with new schema
- ✅ Controller restarts successfully  
- ✅ StatefulSet + Service created
- ✅ Auto-sizing applied (`t3.medium` for 3 replicas)
- ✅ Logs show `"successfully reconciled"`
- ✅ `/metrics` endpoint responds

***

## version control

**Git commit**:

```bash
    git add .
    git commit -m "feat: day2 production mysqlcluster api + controller

    - Custom MySQLClusterSpec with CloudConfig (EC2 instance sizing)
    - StatefulSet controller with resource calculation (t3.micro → m5.large)
    - Geo-aware defaults (nairobi affinity, timezone)
    - Production manager with /metrics /healthz /readyz endpoints
    - Idempotent reconcileResource() with Strategic Merge Patch"
    git push
```

</div>


[def]: ./2-set_up_env.md