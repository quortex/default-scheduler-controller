# default-scheduler-controller

## Overview
This project is a controller for Kubernetes allowing to override default-scheduler at runtime.

Useful in fairly specific situations, this tool is intended to be installed on clusters with several schedulers in order to use an alternative default scheduler.

The most common use case is the use of an alternative scheduler in a managed cluster where there is no possibility to replace the default scheduler.

## Prerequisites

### Kubernetes
A [Kubernetes](https://kubernetes.io/) cluster of version v1.16 or later is required.

### CertManager
This operator needs [CertManager](https://cert-manager.io/docs/) v0.11 or later to be installed on the cluster.


## Installation

### Helm

Follow default-scheduler-controller documentation for Helm deployment [here](./helm/default-scheduler-controller).


## License
Distributed under the Apache 2.0 License. See `LICENSE` for more information.

## Versioning
We use [SemVer](http://semver.org/) for versioning.
