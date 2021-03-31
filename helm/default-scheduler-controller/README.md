# default-scheduler-controller

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.1.0](https://img.shields.io/badge/AppVersion-0.1.0-informational?style=flat-square)

A controller responsible of pods default scheduler overriding.

## Prerequisites

### Kubernetes
A [Kubernetes](https://kubernetes.io/) cluster of version v1.16 or later is required.

### CertManager
This operator needs [CertManager](https://cert-manager.io/docs/) v0.11 or later to be installed on the cluster.

## Usage

In a kubernetes cluster, by default the schedulerName attribute of a pod is "default-scheduler".
This tool handles a mutating webhook to override the defaults-scheduler attribute for new pods created.

## Installation

1. Add default-scheduler-controller helm repository

```sh
helm repo add default-scheduler-controller https://quortex.github.io/default-scheduler-controller
```

2. Create a namespace for default-scheduler-controller

```sh
kubectl create ns default-scheduler-controller-system
```

3. Deploy the appropriate release.

```sh
helm install default-scheduler-controller default-scheduler-controller/default-scheduler-controller -n default-scheduler-controller-system
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| manager.image.repository | string | `"eu.gcr.io/quortex-registry-public/default-scheduler-controller"` | default-scheduler-controller manager image repository. |
| manager.image.tag | string | `"latest"` | default-scheduler-controller manager image tag. |
| manager.image.pullPolicy | string | `"IfNotPresent"` | default-scheduler-controller manager image pull policy. |
| manager.resources | object | `{}` | default-scheduler-controller manager container required resources. |
| manager.schedulerName | string | `"default-scheduler"` | The scheduler to set as default scheduler for all pods. |
| kubeRBACProxy.enabled | bool | `true` |  |
| kubeRBACProxy.image.repository | string | `"gcr.io/kubebuilder/kube-rbac-proxy"` | kube-rbac-proxy image repository. |
| kubeRBACProxy.image.tag | string | `"v0.5.0"` | kube-rbac-proxy image tag. |
| kubeRBACProxy.image.pullPolicy | string | `"IfNotPresent"` | kube-rbac-proxy image pull policy. |
| kubeRBACProxy.resources | object | `{}` | kube-rbac-proxy container required resources. |
| replicaCount | int | `1` | Number of desired pods. |
| imagePullSecrets | list | `[]` | A list of secrets used to pull containers images. |
| nameOverride | string | `""` | Helm's name computing override. |
| fullnameOverride | string | `""` | Helm's fullname computing override. |
| podAnnotations | object | `{}` | Annotations to be added to pods. |
| deploymentAnnotations | object | `{}` | Annotations to be added to deployment. |
| terminationGracePeriod | int | `30` | How long to wait for pods to stop gracefully. |
| nodeSelector | object | `{}` | Node labels for default-scheduler-controller pod assignment. |
| tolerations | list | `[]` | Node tolerations for default-scheduler-controller scheduling to nodes with taints. |
| affinity | object | `{}` | Affinity for default-scheduler-controller pod assignment. |

