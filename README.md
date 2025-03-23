# hse-wishlist
HSE course project for Industrial Software Engineering



## Тестовый прогон кубера

## Запуск в кубере

> minikube delete

> minikube start

## Запуск манифестов в кубере

> kubectl apply -f deployments/k8s/postgres-pvc.yaml

> kubectl apply -f deployments/k8s/postgres-deployment.yaml

> kubectl apply -f deployments/k8s/postgres-service.yaml

> kubectl apply -f deployments/k8s/tasks-svc-configmap.yaml

> kubectl apply -f deployments/k8s/migrate-configmap.yaml

> kubectl apply -f deployments/k8s/migrate-up-job.yaml

> kubectl apply -f deployments/k8s/tasks-svc-deployment.yaml

> kubectl apply -f deployments/k8s/tasks-svc-service.yaml

> kubectl apply -f deployments/k8s/tasks-svc-ingress.yaml

> kubectl apply -f deployments/k8s/fluent-bit-configmap.yaml

> kubectl apply -f deployments/k8s/fluent-bit-daemonset.yaml

> minikube status

## Логи:

> kubectl get pods

> kubectl logs -f (pod name)

> kubectl logs job/migrate

### Чтобы сделать туннель для Postgres
> kubectl port-forward svc/postgres 5432:5432

### Чтобы сделать туннель для Приложения
> kubectl port-forward svc/tasks-svc 8082:8082

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

#### Jobs:
> kubectl delete -f deployments/k8s/migrate-up-job.yaml

> kubectl delete -f deployments/k8s/migrate-down-job.yaml

> kubectl delete -f deployments/k8s/migrate-configmap.yaml

#### Persistent Volume:
> kubectl delete -f deployments/k8s/postgres-pvc.yaml

#### Postgres
> kubectl delete -f deployments/k8s/postgres-deployment.yaml

> kubectl delete -f deployments/k8s/postgres-service.yaml

#### App
> kubectl delete -f deployments/k8s/tasks-svc-configmap.yaml

> kubectl delete -f deployments/k8s/tasks-svc-deployment.yaml

> kubectl delete -f deployments/k8s/tasks-svc-service.yaml

> kubectl delete -f deployments/k8s/tasks-svc-ingress.yaml

#### Fluent-Bit
> kubectl delete -f deployments/k8s/fluent-bit-configmap.yaml

> kubectl delete -f deployments/k8s/fluent-bit-daemonset.yaml

### Зачистка ресурсов:
> kubectl delete deployments --all

> kubectl delete services --all

> kubectl delete jobs --all

> kubectl delete configmaps --all

> kubectl delete daemonsets --all

> kubectl delete pvc --all
