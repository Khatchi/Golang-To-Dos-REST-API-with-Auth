# Load environment variables
-include .env
export

# Fallback DB URL
DATABASE_URL ?= postgres://postgres:kachi@localhost:5432/todo_api_yt?sslmode=disable

# Migration directory
MIGRATIONS_PATH = ./migrations

.PHONY: migrate-create migrate-up migrate-down migrate-down-all migrate-version migrate-force help

# Create a new migration
migrate-create:
	@echo "Creating migration..."
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

# Run all up migrations
migrate-up:
	@echo "Running migrations..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

# Rollback last migration
migrate-down:
	@echo "Rolling back last migration..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

# Rollback ALL migrations
migrate-down-all:
	@echo "Rolling back all migrations..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down

# Show current migration version
migrate-version:
	@echo "Migration version:"
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

# Force version (fix dirty database)
migrate-force:
	@echo "Forcing migration version..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(version)

# Help
help:
	@echo "Available commands:"
	@echo "  make migrate-create name=create_table_name"
	@echo "  make migrate-up"
	@echo "  make migrate-down"
	@echo "  make migrate-down-all"
	@echo "  make migrate-version"
	@echo "  make migrate-force version=1"