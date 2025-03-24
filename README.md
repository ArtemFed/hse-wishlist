# hse-wishlist

HSE course project for Industrial Software Engineering

## Тестовый прогон кубера с фотками

### Блок 1
1. Приложение имеет 4 HTTP API ручки

![Screenshot 2025-03-24 at 03.42.32.png](images%2FScreenshot%202025-03-24%20at%2003.42.32.png)

Получили все аккунты и запомнили id админа (мы сейчас под его токеном)
![Screenshot 2025-03-24 at 03.29.56.png](images%2FScreenshot%202025-03-24%20at%2003.29.56.png)

Создаём таску:
![Screenshot 2025-03-24 at 03.31.21.png](images%2FScreenshot%202025-03-24%20at%2003.31.21.png)

Таска есть в списке и проставлено что её сделал админ, хотя эта инфа была только в токене (передал через контекст):
![Screenshot 2025-03-24 at 03.32.26.png](images%2FScreenshot%202025-03-24%20at%2003.32.26.png)

3. Приложение поддерживает авторизацию

Login:
![Screenshot 2025-03-24 at 03.28.29.png](images%2FScreenshot%202025-03-24%20at%2003.28.29.png)

Любая другая ручка без Authorization Header
![Screenshot 2025-03-24 at 03.29.09.png](images%2FScreenshot%202025-03-24%20at%2003.29.09.png)

Любая другая ручка с "чужим" JWT токеном
![Screenshot 2025-03-24 at 03.29.38.png](images%2FScreenshot%202025-03-24%20at%2003.29.38.png)

Ручка с корректным токеном
![Screenshot 2025-03-24 at 03.29.56.png](images%2FScreenshot%202025-03-24%20at%2003.29.56.png)

3. API покрыто тестами

![Screenshot 2025-03-24 at 03.36.35.png](images%2FScreenshot%202025-03-24%20at%2003.36.35.png)

4. Внешняя БД для хранения пользователей и бизнесинформации
(тоннель к куберу)

![Screenshot 2025-03-24 at 03.27.45.png](images%2FScreenshot%202025-03-24%20at%2003.27.45.png)

5. Схема базы данных создаётся при запуске

(вот есть job'ы миграций (три штуки) на up и down)
![Screenshot 2025-03-24 at 03.21.36.png](images%2FScreenshot%202025-03-24%20at%2003.21.36.png)

6. Схема базы данных отражается в код при сборке. Несоответствие ORM-моделей, запросов и схемы приводит к ошибке сборки.

Такое ORM в клубе Go осуждается, а я религию не нарушаю. А так из-за динамических фильтрационных запросов у меня бы не получилось сгенерить через `https://github.com/sqlc-dev/sqlc`.
Я попробовал - вышло очень страшно с либой - удалил.

### Блок 2

1. Приложение поддерживает логирование.

![Screenshot 2025-03-24 at 03.52.18.png](images%2FScreenshot%202025-03-24%20at%2003.52.18.png)

2. Приложение поддерживает метрики.

Grafana + Prometheus пишутся (использовал стандартный Go Дашборд)
![Screenshot 2025-03-24 at 03.23.20.png](images%2FScreenshot%202025-03-24%20at%2003.23.20.png)

Очень хотел завести Jaeger на OpenTelemetry, но кубер не съел(

3. Приложение может быть запущено в Kubernetes (приложение, БД, логирование,
   балансировщик и сервис).

![Screenshot 2025-03-24 at 03.21.24.png](images%2FScreenshot%202025-03-24%20at%2003.21.24.png)

![Screenshot 2025-03-24 at 03.54.18.png](images%2FScreenshot%202025-03-24%20at%2003.54.18.png)

Всё падает и встаёт само.

4. Поддерживается сборка логов приложения и всех взаимодействующих с приложением инфраструктурных объектов в Kubernetes.

![Screenshot 2025-03-24 at 03.21.55.png](images%2FScreenshot%202025-03-24%20at%2003.21.55.png)
Сбору сделал, но в фильтрах запутался(

### Блок 3

1. Каждый коммит ... (плашка на GitHub).

Когда плохо:
![Screenshot 2025-03-24 at 03.37.19.png](images%2FScreenshot%202025-03-24%20at%2003.37.19.png)

Когда хорошо:
![Screenshot 2025-03-24 at 03.37.28.png](images%2FScreenshot%202025-03-24%20at%2003.37.28.png)

2. По всем API-методам есть Swagger-документация,

Будет по такой ссылке:
http://localhost:8082/api/v1/swagger-ui/index.html#/Auth/post_auth

![Screenshot 2025-03-24 at 03.25.18.png](images%2FScreenshot%202025-03-24%20at%2003.25.18.png)

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

Я очень старался, но otel не заработал((

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

> kubectl delete namespace monitoring

> kubectl delete job migrate

> kubectl delete job migrate-down

> kubectl delete pvc postgres-pvc

> kubectl delete configmap fluent-bit-config

> kubectl delete daemonset fluent-bit

> kubectl delete configmap otel-config

> kubectl delete deployment otel-collector

> kubectl delete service otel-collector

> kubectl delete configmap prometheus-config

> kubectl delete deployment prometheus

> kubectl delete service prometheus

> kubectl delete configmap grafana-datasources

> kubectl delete deployment grafana

> kubectl delete service grafana

> kubectl delete configmap jaeger-config

> kubectl delete deployment jaeger

> kubectl delete service jaeger

### Зачистка ресурсов:

> kubectl delete deployments --all

> kubectl delete services --all

> kubectl delete jobs --all

> kubectl delete configmaps --all

> kubectl delete daemonsets --all

> kubectl delete pvc --all
