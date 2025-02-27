---
title: "Installing with Helm (recommended)"
linkTitle: "Installing with Helm (recommended)"
weight: 10
---

## Installing 

The Aiven Operator for Kubernetes can be installed via [Helm](https://helm.sh/). 

First add the [Aiven Helm repository](https://github.com/aiven/aiven-charts):

```shell
helm repo add aiven https://aiven.github.io/aiven-charts && helm repo update
```

### Installing Custom Resource Definitions

```shell
helm install aiven-operator-crds aiven/aiven-operator-crds
```

Verify the installation:
```shell
kubectl api-resources --api-group=aiven.io
```
The output is similar to the following:

```{ .shell .no-copy }
NAME                  SHORTNAMES   APIVERSION          NAMESPACED   KIND
connectionpools                    aiven.io/v1alpha1   true         ConnectionPool
databases                          aiven.io/v1alpha1   true         Database
... < several omitted lines >
```

### Installing the Operator

```shell
helm install aiven-operator aiven/aiven-operator
```

!!! note
    Installation will fail if webhooks are enabled and the CRDs for the cert-manager are not installed.

Verify the installation: 
```shell
helm status aiven-operator
```

The output is similar to the following:

```{ .shell .no-copy }
NAME: aiven-operator
LAST DEPLOYED: Fri Sep 10 15:23:26 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

It is also possible to install the operator without webhooks enabled:
```shell
helm install aiven-operator aiven/aiven-operator --set webhooks.enabled=false
```

### Configuration Options

Please refer to the [values.yaml](https://github.com/aiven/aiven-charts/blob/main/charts/aiven-operator/values.yaml) of the chart.

#### Installing without full cluster administrator access
There can be some scenarios where the individual installing the Helm chart does not have the ability to provision cluster-wide resources (e.g. ClusterRoles/ClusterRoleBindings). In this scenario, you can have a cluster administrator manually install the [ClusterRole](../../../charts/aiven-operator/templates/cluster_role.yaml) and [ClusterRoleBinding](../../../charts/aiven-operator/templates/cluster_role_binding.yaml) the operator requires prior to installing the Helm chart specifying `false` for the `clusterRole.create` attribute. 

## Uninstalling 

!!! important
    Please see [this page](uninstalling.md) for more information.

Find out the name of your deployment:

```shell
helm list
```

The output has the name of each deployment similar to the following: 

```{ .shell .no-copy }
NAME               	NAMESPACE	REVISION	UPDATED                                 	STATUS  	CHART                     	APP VERSION
aiven-operator     	default  	1       	2021-09-09 10:56:14.623700249 +0200 CEST	deployed	aiven-operator-v0.1.0     	v0.1.0     
aiven-operator-crds	default  	1       	2021-09-09 10:56:05.736411868 +0200 CEST	deployed	aiven-operator-crds-v0.1.0	v0.1.0
```

Remove the CRDs:

```shell
helm uninstall aiven-operator-crds
```

The confirmation message is similar to the following:

```{ .shell .no-copy }
release "aiven-operator-crds" uninstalled
```

Remove the operator:

```shell
helm uninstall aiven-operator
```

The confirmation message is similar to the following:

```{ .shell .no-copy }
release "aiven-operator" uninstalled
```

