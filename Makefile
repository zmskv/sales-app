ifneq (,$(wildcard .env))
    include .env
    export
endif

MIGRATE=migrate -path ./migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)"

.PHONY: app-start migrate-up migrate-down build

app-start:
	go run cmd/main.go

build:
	go build github.com/zmskv/sales-app/cmd

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down
