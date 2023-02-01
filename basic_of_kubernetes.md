# Basic of Kubernetes (K8s)
![](https://badgen.net/badge/status/in%20progress/orange) ![](https://badgen.net/badge/version/v0.0.1/cyan)

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

If you have many applications on your server, it will take a long time to maintain, deploy, and monitor all of your applications one by one. With Kubernetes, they will help you to maintain, deploy, and monitor easily helped by a system. Because of that, Kubernetes was born! 😄️ Kubernetes will help you with their framework for maintaining etc your applications easily with smaller mistakes.

## Kubernetes

### Kubernetes Component

### Kubernetes API

### Kubernetes Object

* Nodes
* Pods
* Deployment
* Service
* Ingress
* Volumes, Persistent Volumes & Persistent Volumes Claim
* Config Maps
* Screts

## Reference

Title | URL
--- | ---
Kubernetes Documentation | https://kubernetes.io/docs/home
Kubernetes Documentation Reference | https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#-strong-api-overview-strong
NGINX Ingress Controller | https://kubernetes.github.io/ingress-nginx/deploy
ChatGPT | https://chat.openai.com/
