# Kubernetes Loadbalancing test

## Start Microk8

```bash
microk8s status --wait-ready
```

### Start Dashboard

```bash
microk8s dashboard-proxy
```

## Build Container Image & import into Microk8

```bash
bash build
```

## Start Deployment, Service, Scaler

```bash
microk8s kubectl apply -f ./webapp-deployment.yaml
microk8s kubectl apply -f ./webapp-service.yaml
microk8s kubectl apply -f ./webapp-scaler.yaml
```

## Scale
```bash
microk8s kubectl scale -n default deployment webapp-deployment --replicas=2
```

## Restart Deployment (update)
```bash
kubectl rollout restart -f ./webapp-deployment.yaml
```

## Delete Deployment, Service, Scaler

```bash
microk8s kubectl delete -f ./webapp-deployment.yaml
microk8s kubectl delete -f ./webapp-service.yaml
microk8s kubectl delete -f ./webapp-scaler.yaml
```
