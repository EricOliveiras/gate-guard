# Variables
DB_DRIVER = postgres
DB_USER = admin
DB_PASS = admin
DB_NAME = gate-guard-db
DB_HOST = localhost
DB_PORT = 5432

APP_NAME = gate-guard
DB_MIGRATIONS_DIR = internal/migrations

DB_CONN_STR = "${DB_DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Commands
.PHONY: all build run migrate-up migrate-down

all: build run

build:
	@echo "Building the Go application..."
	@go build -o gate-guard cmd/main.go
	@chmod +x gate-guard

go-test:
	@echo "Running tests..."
	@go test test/*_test.go

run:
	@echo "Running the Go application..."
	@./$(APP_NAME)

migrate-up:
	@echo "Applying database migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR) up

migrate-down:
	@echo "Reverting database migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR) down

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

.PHONY: clean
