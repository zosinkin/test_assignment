include .env
export

export PROJECT_ROOT=${shell pwd}

# Запуск контейнера PostgreSQL
env-up:
	@docker compose up -d subscription-postgres

# Остановка контейнера PostgreSQL
env-down:
	@docker compose down subscription-postgres

# Полная очистка окружения (удаляет volume с данными БД)
env-cleanup:
	@read -p "Should I clean all volume files in the environment? There is a risk of data loss! [y/N]:" ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down subscription-postgres port-forwarder && \
		rm -rf out/pgdata && \
		echo "The environment files have been cleaned"; \
	else \
		echo "The environment cleanup has been canceled"; \
	fi

logs-cleanup: ## env: Очистить файлы логов из out/logs
	@read -p "Очистить все log файлы? Опасность утери логов. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Файлы логов очищены"; \
	else \
		echo "Очистка логов отменена"; \
	fi

# Показать статус контейнеров
ps: 
	@docker compose ps

# Создание новой миграции (нужно передать seq=название)
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm subscription-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

# Применить все миграции вверх
migrate-up:
	@make migrate-action action=up

# Откатить миграции
migrate-down:
	@make migrate-action action=down

# Универсальная команда для выполнения миграций (up/down)
migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm subscription-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@subscription-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

# Генерация swagger-документации
swagger-gen: 
	@docker compose run --rm swagger \
		init \
		-g cmd/app/main.go \
		-o docs \
		--parseInternal \
		--parseDependency

# Проброс порта PostgreSQL на localhost:5432
env-port-forward:
	@docker compose up -d port-forwarder

# Остановить проброс порта
env-port-close:
	@docker compose down port-forwarder

## Golang приложение: Запустить локально
app-run: 
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/app/main.go

## Golang приложение: Запустить в Docker                     
app-deploy: 
	@docker compose up -d --build app


## Golang приложение: Остановить в Docker
app-deloy-down:
	@docker compose down app