# kubectl commands

## Help

```bash
# List api resources
$ kubectl api-resources

$ kubectl explain pods
$ kubectl explain pod.spec

```

## Resource general

```bash

# Create resources
$ kubectl create -f ${file_name}

# Create resources
$ kubectl create -f ${file_name} -n ${namespace_name}

# Delete all resources in a namespace
$ kubectl delete all --all

# Edit resources
$ kubectl edit ${resource} ${resource_name}
```

## Node

```bash

# List nodes
$ kubectl get node

```

## Pods - Basic

```bash

# Full yaml of a deployed pod
$ kubectl get pods ${pod_name} -o yaml
# Full json of a deployed pod
$ kubectl get pods ${pod_name} -o json

# Describe pod
$ kubectl describe pod ${pod_name}

# List pods
$ kubectl get pods

# List pods with labels
$ kubectl get pods --show-labels

# List pods details
$ kubectl get pods -o wide

# List pods with label as column
$ kubectl get pods -L ${labels_with_comma_separated}
e.g. $ kubectl get pods -L creaion_method,env

# List pods with namespaces
$ kubectl get pods --namespace ${namespace_name}
$ kubectl get pods -n ${namespace_name}

# Delete pods by name
$ kubectl delete pods ${pod_name}
$ kubectl delete pods ${pod_name_1} ${pod_name_2} ...

# Delete pods using label selectors
$ kubectl delete pods -l ${label_key}=${label_value}
e.g. $ kubectl delete pods -l type=web

# Delete pods by deleting the whole namespace
$ kubectl delete ns ${namespace_name}

# Delete all pods in a namespace, while keeping the namespace
$ kubectl delete pods --all
```

## Labels

```bash

# Add label to existing pod
$ kubectl label pods ${pod_name} ${label_key}=${label_value}
e.g. $ kubectl label po my-pod hello=world

# Update label to existing pod
$ kubectl label pods ${pod_name} ${labem_key}=${label_value} --overwrite

# List all pods include label
$ kubectl get pods -l ${label}

# List all pods do not include label
#   only single quote is allowed
$ kubectl get pods -l '!${label}'

```

## Annotations

```bash

# Add annotation
$ kubectl annotate pod ${pod_name} ${annotation_key}=${annotation_value}
e.g. $ kubectl annotate pod my-pod mycompany.com/sampleannotation="hello world"

```

## Namespaces

```bash

# List all namespaces
$ kubectl get namespaces
$ kubectl get ns

# Create namespace
$ kubectl create namespace ${namespace_name}

```

## Log

```bash

# Retrieving a pod's log
$ kubectl logs ${pod_name}

# Retrieving container's log

$ kubectl logs ${pod_name} -c ${container_name}

```

## Resource - ReplicationController(rc), ReplicaSet(rs)

```bash

# Get
$ kubectl get rs
$ kubectl get rc

# Desribe
$ kubectl describe rs

# Scale replica
$ kubectl scale rs ${pod_name} --replicas=${count}
$ kubectl scale rc ${pod_name} --replicas=${count}

# Delete
$ kubectl delete rs ${replica_set_name}
$ kubectl delete rc ${replication_controller_name}

# Delete, keep its pods running
$ kubectl delete rs ${replication_controller_name} --cascade=false
$ kubectl delete rc ${replication_controller_name} --cascade=false
```

## Resource - Job

```bash

# Seeing a Job
$ kubectl get jobs
$ kubectl get pods

# Scaling a job
$ kubectl scale job ${job_anme} --replicas ${parallelism}
```

## Others

```bash

# Forwarding a local network port to a port in the pod
$ kubectl port-forward ${pod_name} ${machine_port}:${pod_port}
e.g. $ kubectl port-forward my-pod 8888:8080

```
