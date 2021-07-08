# Kubernetes deployment

# Install App

## 1. Deployment

Deploy assistant to your Kubernetes Cluster from zero.

### Create a namespace

Create a namespace in the cluster to manage all the resources of Xconf. Edit `namespace.yaml` if you need.

```kubectl create namespace.yaml```

### Edit config files

### Create ConfigMap

```
kubectl create -f configmaps.yaml
```

### Create Kubernets resources

```
kubectl create -f gateway.yaml
```

### Access gateway

By default the gateway will be exposed by LoadBalancer, check service `gateway` under namespace `assistant` for detail.

`kubectl describe --namespace assistant service gateway`

You will find the Ingress address and the port.

Use url `http://addr:port/` to access the web.

## 2. Update

Update the Kubernets configuration if resources have already existed.

Modify the config file corresponding to you needs.

Apply the update:

```kubectl apply -f FILENAME```

## 3. Remove

You can remove all the deployments by deleting the whole namespace.

```kubectl delete namespace assistant```
