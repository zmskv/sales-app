# Загружаем переменные из .env 
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Команда для миграций
MIGRATE=migrate -path ./migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)"

.PHONY: app-start migrate-up migrate-down

app-start:
	go run cmd/main.go

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down
