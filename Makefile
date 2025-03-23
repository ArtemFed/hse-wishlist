include .env

gen:
	oapi-codegen --config ./services/tasks/.codegen/task-codegen-config.yaml ./services/tasks/.codegen/task-codegen.yaml

# Взлёты

up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-all:
	docker compose -f ./deployments/compose.yaml up -d --build
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

# Падения

down:
	docker compose down tasks-svc

# Observability

up-obs:
	docker compose -f ./deployments/compose.yaml up -d --build

down-obs:
	docker compose -f ./deployments/compose.yaml down

# Миграции

migrate-up:
	migrate -path ./services/tasks/migrations -database 'postgres://$(TASKS_POSTGRES_USER):$(TASKS_POSTGRES_PASSWORD)@$(TASKS_POSTGRES_HOST_LOCAL):$(TASKS_POSTGRES_PORT_EXTERNAL)/$(TASKS_POSTGRES_NAME)?sslmode=disable' up

migrate-down:
	migrate -path ./services/tasks/migrations -database 'postgres://$(TASKS_POSTGRES_USER):$(TASKS_POSTGRES_PASSWORD)@$(TASKS_POSTGRES_HOST_LOCAL):$(TASKS_POSTGRES_PORT_EXTERNAL)/$(TASKS_POSTGRES_NAME)?sslmode=disable' down 1
