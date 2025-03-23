# hse-wishlist
HSE course project for Industrial Software Engineering



## Тестовый прогон кубера

kubectl create deployment k8s-hse-wishlist --image=kicbase/echo-server:1.0
kubectl expose deployment k8s-hse-wishlist --type=LoadBalancer --port=8082
kubectl get service

## Запуск в кубере

> minikube delete

> minikube start

> kubectl apply -f deployments/k8s/postgres-pvc.yaml
> kubectl apply -f deployments/k8s/postgres-deployment.yaml
> kubectl apply -f deployments/k8s/postgres-service.yaml
> kubectl apply -f deployments/k8s/tasks-svc-configmap.yaml
> kubectl apply -f deployments/k8s/migrate-configmap.yaml
> kubectl apply -f deployments/k8s/migrate-job.yaml
> kubectl apply -f deployments/k8s/tasks-svc-deployment.yaml
> kubectl apply -f deployments/k8s/tasks-svc-service.yaml
> kubectl apply -f deployments/k8s/tasks-svc-ingress.yaml
> kubectl apply -f deployments/k8s/fluent-bit-configmap.yaml
> kubectl apply -f deployments/k8s/fluent-bit-daemonset.yaml

> minikube status

> kubectl get pods

> kubectl logs -f <pod name>

Чтобы сделать туннель для Postgres
> kubectl port-forward svc/postgres 5432:5432

Чтобы сделать туннель для Приложения
> kubectl port-forward svc/tasks-svc 8082:8082


