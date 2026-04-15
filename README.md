# 🚀 Запуск приложения через Makefile

## Основные сценарии запуска

Есть два способа запуска приложения:

- локально через Go
- в Docker

В обоих случаях сначала нужно поднять PostgreSQL и применить миграции.

---

## ⚡ Быстрый старт

### Локальный запуск приложения

```bash
make env-up
make env-port-forward
make migrate-up
swagger-gen
make app-run
```

### Запуск через Docker
```bash
make env-up
make env-port-forward
make migrate-up
swagger-gen
make app-deploy
```

После этого приложение будет доступно по адресу:

`http://localhost:5050`

Swagger:
`http://localhost:5050/swagger`

За дополнительными командами обращаться в `Makefile`
