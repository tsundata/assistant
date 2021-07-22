# Kubernetes deployment

# Install App

## 1. Deployment

Deploy assistant to your Kubernetes Cluster from zero.

### Create a namespace

Create a namespace in the cluster to manage all the resources of Assistant. Edit `namespace.yaml` if you need.

```kubectl create namespace.yaml```

### Edit config files

### Create ConfigMap

```
kubectl create -f configmaps.yaml
```

### Create Kubernets resources

```
kubectl create -f ./app
```

### Access gateway/web

By default the gateway will be exposed by LoadBalancer, check service `gateway/web` under namespace `assistant` for detail.

`kubectl describe --namespace assistant service gateway`
`kubectl describe --namespace assistant service web`

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
