# BYO kubernetes operators with kubebuilder and Go

## flow
* theory of controllers
    * Introduction & Prerequisites
    * What is a Controller? (The Observe-Compare-Act Loop)
    * Idempotency in Controllers
    * Deep Dive: The Reconcile Loop (Happy Path, Sad Path, & Error Handling)
    * The Foundation of Writing Operators
    * What is an Operator? (The "Helper" Analogy)
    * CRDs (Custom Resource Definitions) and CRs (Custom Resources)
* kubernetes extensibility
    * Kubernetes as an SDK & Extensibility
    * Networking, Storage, & Admission Controllers
    * Internal Developer Platforms (IDP) & Platform Engineering
    * Bootstrapping with Kubebuilder
* set up the env
    * Setting up the Local Environment (K3D, Docker)
    * Introduction to the Kubebuilder Framework
    * Project Initialisation (kubebuilder init)
    * Exploring Scaffolding (Makefiles, Dockerfiles, main.go)
* build API and logic
    * Creating your first API (kubebuilder create api)
    * Defining EC2 Instance Types & Specs in Go
    * Understanding TypeMeta and ObjectMeta
    * Internal Controller Logic Breakdown
    * Deep Dive: Manager Architecture & Controller-Runtime
    * Cert Watchers, Health Checks, & Prometheus Metrics
    * Initialising the Manager in main.go
* hands-on dev
    * Implementing the Reconcile Loop Logic
    * Custom Resource Definitions (CRDs) in Action
    * Running the Operator Locally
    * AWS SDK Integration in Go
    * Using Finalizers for Cleanup Logic
    * Creating EC2 Instances on AWS via the Operator
    * Implementing Waiters for Instance State (Running/Terminated)
* advanced internals & dev
    * Idempotency & Reconciler Loop Internals
    * How Informers, Caches, and WorkQueues Work
    * Handling Object Deletion & Timestamps
    * Packaging the Operator with Helm
    * Deploying to Kubernetes (RBAC & Service Accounts)
    * Conclusion & Future Steps
