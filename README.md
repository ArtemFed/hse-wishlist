# hse-wishlist
HSE course project for Industrial Software Engineering



## Тестовый прогон кубера

## Запуск в кубере

> minikube delete

> minikube start

> kubectl create namespace monitoring 

## Запуск манифестов в кубере

> kubectl apply -f deployments/k8s/postgres-pvc.yaml

> kubectl apply -f deployments/k8s/postgres-deployment.yaml

> kubectl apply -f deployments/k8s/postgres-service.yaml

> kubectl apply -f deployments/k8s/tasks-svc-configmap.yaml

> kubectl apply -f deployments/k8s/migrate-configmap.yaml

> kubectl apply -f deployments/k8s/migrate-up-job.yaml

> kubectl apply -f deployments/k8s/migrate-down-job.yaml

> kubectl apply -f deployments/k8s/tasks-svc-deployment.yaml

> kubectl apply -f deployments/k8s/tasks-svc-service.yaml

> kubectl apply -f deployments/k8s/tasks-svc-ingress.yaml

> kubectl apply -f deployments/k8s/fluent-bit-configmap.yaml

> kubectl apply -f deployments/k8s/fluent-bit-daemonset.yaml

> kubectl apply -f deployments/k8s/prometheus-configmap.yaml

> kubectl apply -f deployments/k8s/prometheus-deployment.yaml

> kubectl apply -f deployments/k8s/prometheus-service.yaml

> kubectl apply -f deployments/k8s/prometheus-ingress.yaml

> kubectl apply -f deployments/k8s/grafana-configmap.yaml

> kubectl apply -f deployments/k8s/grafana-deployment.yaml

> kubectl apply -f deployments/k8s/grafana-service.yaml

> kubectl apply -f deployments/k8s/grafana-ingress.yaml

> kubectl apply -f deployments/k8s/jaeger-deployment.yaml

> kubectl apply -f deployments/k8s/jaeger-service.yaml

> kubectl apply -f deployments/k8s/otel-configmap.yaml

> kubectl apply -f deployments/k8s/otel-deployment.yaml

> kubectl apply -f deployments/k8s/otel-service.yaml

> minikube status

## Логи:

> kubectl get pods

> kubectl get pods --namespace=monitoring

> kubectl get pods --all-namespaces

> kubectl logs -f (pod name)

> kubectl logs job/migrate

### Делаем туннели (хотя с таким количеством надо было уже в DNS запихать)
> kubectl port-forward svc/postgres 5432:5432

> kubectl port-forward svc/tasks-svc 8082:8082

> kubectl port-forward -n monitoring svc/prometheus 9090:9090

> kubectl port-forward -n monitoring svc/grafana 3000:3000

> kubectl port-forward -n monitoring svc/jaeger 16686:16686

## Работа с приложением:

Swagger схема:
http://localhost:8082/api/v1/swagger-ui/index.html#/Auth/post_auth

```
Body: {
    "login": "admin",
    "password": "admin"
}
```

Получаем примерно ответ:
```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI3NzAwNjEsImlhdCI6MTc0Mjc2ODI2MSwidXNlcl91dWlkIjoiZDRlM2YzNzctMjMyMC00Mzk1LWFlYmUtNWFhYTQ0M2U3YzM0IiwibG9naW4iOiJhZG1pbiJ9.mgbavbJzM8aVsa4vwTmkIbcP6eYI3fhRdoOSp-GFPJw"
}
```

Регаем в Swagger UI сверху

Накидаем данных в Create:
http://localhost:8082/api/v1/swagger-ui/index.html#/Accounts/post_accounts

```
{
  "login": "Artem",
  "password": "pass1"
}
```

```
{
  "login": "Oleg",
  "password": "pass2"
}
```

## Зачистка:

### Зачистка манифестов:

> kubectl delete -f deployments/k8s/migrate-up-job.yaml

> kubectl delete -f deployments/k8s/migrate-down-job.yaml

> kubectl delete -f deployments/k8s/migrate-configmap.yaml

> kubectl delete -f deployments/k8s/postgres-pvc.yaml

> kubectl delete -f deployments/k8s/postgres-deployment.yaml

> kubectl delete -f deployments/k8s/postgres-service.yaml

> kubectl delete -f deployments/k8s/tasks-svc-configmap.yaml

> kubectl delete -f deployments/k8s/tasks-svc-deployment.yaml

> kubectl delete -f deployments/k8s/tasks-svc-service.yaml

> kubectl delete -f deployments/k8s/tasks-svc-ingress.yaml

> kubectl delete -f deployments/k8s/fluent-bit-configmap.yaml

> kubectl delete -f deployments/k8s/fluent-bit-daemonset.yaml

### Зачистка ресурсов:
> kubectl delete deployments --all

> kubectl delete services --all

> kubectl delete jobs --all

> kubectl delete configmaps --all

> kubectl delete daemonsets --all

> kubectl delete pvc --all
