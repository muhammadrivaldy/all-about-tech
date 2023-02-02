# Basic of Kubernetes (K8s)

![](https://badgen.net/badge/status/completed/green) ![](https://badgen.net/badge/version/v1.0.0/cyan)

## Overview

Kubernetes, called K8s is a platform open source for automation deployment, scaling and managing many containers. K8s is an abbreviation of "Kubernetes" in which "ubernete" are replaced by "8".

In 2014 Google teams state the project Kubernetes will be set as Open source.

### Let's throwback

![](https://d33wubrfki0l68.cloudfront.net/26a177ede4d7b032362289c6fccd448fc4a91174/eb693/images/docs/container_evolution.svg)

* #### Traditional Deployment

  We don't have any mechanism for creating a bondaries between application inside of physical server. This approach raises a new problem, the problem appear because the server have many applications there. So, when one application used many resources in that server, another application will get impacted because they don't have a capable resource.

* #### Virtualized Deployment

  With this mechanism, we can solve the issue from traditional deployment. Why? Because with virtualized deployment, you can run multiple applications with different machines (called virtual machines). But, it's not easy to set up the physical server rather than container deployment.

* #### Container Deployment

  The mechanism of this initiative is really similar to virtualized deployment. But, Kubernetes more excellence rather than virtualized. Kubernetes is better at doing isolation applications because of that the containers are lighter than VM.

### Why used Kubernetes?

If you have many applications on your server, it will take a long time to maintain, deploy, and monitor all of your applications one by one. With Kubernetes, they will help you to maintain, deploy, and monitor easily helped by a system. Because of that, Kubernetes was born! 😄️

Kubernetes will help you with their framework for maintaining etc your applications easily with smaller mistakes.

## Kubernetes

### Kubernetes Component

When we install Kubernetes, we will get a cluster automatically. In a Kubernetes cluster, you will found several worker machines (called nodes) that running the container. Every cluster at least has one node.

Worker nodes host the pods that are the components of the workload. While the control plant will be responsible to manage nodes & pods.

### Kubernetes API

The main function of the Kubernetes control plant is the API server. The API server will connect the user, every single part of the cluster & the external components can be able to connect with each other. In many cases, you can do an operation with tools like kubectl or kubeadm. Basically, the tools will be generating your request in API format and send it to Kubernetes.

You also be able to request the operation directly with API through REST calls.

### Kubernetes Object

Kubernetes objects are persistent entities in the Kubernetes system. Kubernetes used that entities for mapping out the cluster situation. Especially, Kubernetes will map outs:

* What application (container) are running and on which nodes.
* Resources are available for those applications.
* Policy on how the app behaves. Like how the restart policy, upgrades, and fault tolerance.

If you wanna work with Kubernetes objects (for creating, modifying, or deleting) we must use Kubernetes API. To use it, we can use Kubectl or direct requests with REST API.

* #### Nodes

  Nodes represent the virtual machine inside the cluster. Every node will manage by the control plane and inside of the nodes it has many tools that we need to run the pods. In the real case, we are able to have more than one node in the cluster. Usually, for learning cases, we only have one node in the cluster.

  Creating a node on Kubernetes will depend on the cloud server that we use. The way to create a node on GKE & Amazon EKS will be different. But, in the learning case, we are able to create a node with a Kubernetes object. Please remind in your mind, this is not the best way to do it.

  Example a Kubernetes object of nodes:

  ```
  apiVersion: v1
  kind: Node
  metadata:
    name: node-example
  spec:
    podCIDR: 10.200.0.0/16
    externalID: node-example-id
  ```

* #### Pods

  The pod is the smallest that we are able to deploy and manage inside Kubernetes. We can run one or more containers in one pod with sharing resources.

  Example a Kubernetes object of pods:
  
  ```
  apiVersion: v1
  kind: Pod
  metadata:
    name: nginx
  spec:
    containers:
    - name: nginx
      image: nginx:1.14.2
      ports:
      - containerPort: 80
  ```

* #### Deployments

  Deployment gives a declarative update to the pod and the replica set. Basically, to create and run a pod only needs a "pod object" and if we want to create a replica for a pod we are able to use a "replica set object".

  So, why do we need a "deployment object"? The answer is because "deployment object" has several excellent features that easier the process development, among others:

  - Rolling update
  - Rollback a deployment
  - Scaling a deployment
  - Pausing and resuming a rollout

  Example a Kubernetes object of deployments:

  ```
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: nginx-deployment
    labels:
      app: nginx
  spec:
    replicas: 3
    selector:
      matchLabels:
        app: nginx
    template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
          - containerPort: 80
  ```

* #### Service

  Service is used to expose the pod to be able to access with certain conditions (usually that pattern called microservice). Kubernetes will give an IP Address and a single DNS.

  Why do we need a service? Because it's important if we want to set a replica of the pod. So, if other pods want to access a pod, they can access it from DNS that it has. And, service will share the request with a load balancing feature.

  Example a Kubernetes object of service:

  ```
  apiVersion: v1
  kind: Service
  metadata:
    name: my-service
  spec:
    selector:
      app.kubernetes.io/name: MyApp
    ports:
      - protocol: TCP
        port: 80
        targetPort: 9376
  ```

* #### Ingress

  Ingress has a special job to manage access from outside to service in a cluster. Traffic routing is controlled with a policy from the ingress object.

  This is how the system looks like:

  ![](https://www.spectrocloud.com/static/f5c9c043a2705d286584c0c03d42bf51/5ca22/73-image1.png)

  Example a Kubernetes object of ingress:

  ```
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: minimal-ingress
    annotations:
      nginx.ingress.kubernetes.io/rewrite-target: /
  spec:
    ingressClassName: nginx-example
    rules:
    - http:
        paths:
        - path: /testpath
          pathType: Prefix
          backend:
            service:
              name: test
              port:
                number: 80
  ```

* #### Volumes, Persistent Volumes & Persistent Volumes Claim

  Basically, they have a relation each other. Why?

  Because volume is used for connecting storage with a container inside of the pod. So, the container can save and read all files or folders that are connected. If we don't use persistent volume, the volume will be saved at the pod level.

  Persistent volume is a mechanism to save data permanently. So, the data will be safe if we have some failures at a pod that causes the pod to get restarted. Because persistent volume will save the data at the cluster level.

  Persistent volume claim is used by a pod to request some resource to persistent volume. So, a pod will have storage that aligns with the pod needed.

  ![](https://i0.wp.com/blog.knoldus.com/wp-content/uploads/2021/05/1_keV2VBkCHb7cn_Rib0huYg.png?fit=810%2C516&ssl=1)

  Example a Kubernetes object of volume, persistent volume, & persistent volume claim:

  ```
  apiVersion: v1
  kind: PersistentVolume
  metadata:
    name: my-pv
  spec:
    capacity:
      storage: 1Gi
    accessModes:
      - ReadWriteOnce
    persistentVolumeReclaimPolicy: Retain
    storageClassName: standard
    hostPath:
      path: "/data"
  ---
  apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    name: my-pvc
  spec:
    accessModes:
      - ReadWriteOnce
    resources:
      requests:
        storage: 1Gi
    storageClassName: standard
  ---
  apiVersion: v1
  kind: Pod
  metadata:
    name: my-pod
  spec:
    containers:
      - name: my-container
        image: nginx
        volumeMounts:
          - mountPath: "/data"
            name: my-volume
    volumes:
      - name: my-volume
        persistentVolumeClaim:
          claimName: my-pvc
  ```

* #### Config Maps

  Config map is used to save data config that is non-confidential with a format key-value. The pod is able to get the data config as an environment variable, command-line argument, or configuration file inside of volume.

  Example a Kubernetes object of config map:

  ```
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: game-demo
  data:
    player_initial_lives: "3"
    ui_properties_file_name: "user-interface.properties"
  ```

* #### Screts

  The secret has similar behavior to config map. The difference is only in the goals. The secret will use for confidential data, like a password or secret key.

  Example a Kubernetes object of secret:

  ```
  apiVersion: v1
  kind: Secret
  metadata:
    name: my-secret
  type: Opaque
  data:
    username: YWRtaW4=
    password: MWYyZDFlMmU2N2Rm
  ```

## Reference

Title | URL
--- | ---
Kubernetes Documentation | https://kubernetes.io/docs/home
Kubernetes Documentation Reference | https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#-strong-api-overview-strong
NGINX Ingress Controller | https://kubernetes.github.io/ingress-nginx/deploy
ChatGPT | https://chat.openai.com/

## Full Implementation
I created full implementation Kubernetes with Golang and you can take a look at this repository https://github.com/muhammadrivaldy/basic-k8s
