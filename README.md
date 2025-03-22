# hse-wishlist
HSE course project for Industrial Software Engineering



## Тестовый прогон кубера

kubectl create deployment k8s-hse-wishlist --image=kicbase/echo-server:1.0
kubectl expose deployment k8s-hse-wishlist --type=LoadBalancer --port=8080
kubectl get service

## Запуск в кубере

> minikube delete

> minikube start

> minikube status

> kubectl create deployment k8s-hse-wishlist --image=sirdaukar/hse_wishlist:latest

> kubectl expose deployment k8s-hse-wishlist --port=5432

> kubectl get service ***

> kubectl apply -f deployments/k8s/k8s-dpl-tasks-svc-config.yaml

> kubectl apply -f deployments/k8s/k8s-dpl-postgres-config.yaml
