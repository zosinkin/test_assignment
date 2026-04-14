include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d subscription-postgres

env-down:
	@docker compose down subscription-postgres

env-cleanup:
	@read -p "Should I clean all volume files in the environment? There is a risk of data loss! [y/N]:" ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down subscription-postgres port-forwarder && \
		rm -rf out/pgdata && \
		echo "The environment files have been cleaned"; \
	else \
		echo "The environment cleanup has been canceled"; \
	fi

migarte-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутсвует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm subscription-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутсвует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm subscription-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@subscription-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

app-run:
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/app/main.go