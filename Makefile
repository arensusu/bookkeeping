include .env

.PHONY: migrate-up migrate-down

migrate-up:
	migrate -path database/migration -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" -verbose down

