# Variables
DB_DRIVER = postgres
DB_USER = admin
DB_PASS = admin
DB_NAME = gate-guard-db
DB_HOST = localhost
DB_PORT = 5432

# Test Variables
DB_TEST_DRIVER = postgres
DB_TEST_USER = admin
DB_TEST_PASS = admin
DB_TEST_NAME = gate-guard-test-db
DB_TEST_HOST = localhost
DB_TEST_PORT = 5433

APP_NAME = gate-guard
DB_MIGRATIONS_DIR = internal/migrations

DB_CONN_STR = "${DB_DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
DB_CONN_STR_TEST = "${DB_TEST_DRIVER}://${DB_TEST_USER}:${DB_TEST_PASS}@${DB_TEST_HOST}:${DB_TEST_PORT}/${DB_TEST_NAME}?sslmode=disable"

# Commands
.PHONY: all build run migrate-up migrate-down

all: build run

build:
	@echo "Building the Go application..."
	@go build -o gate-guard cmd/main.go
	@chmod +x gate-guard

go-test:
	@make migrate-up-test
	@echo "Running tests..."
	@go test test/*_test.go -v
	@yes | make migrate-down-test

run:
	@echo "Running the Go application..."
	@./$(APP_NAME)

migrate-up:
	@echo "Applying database migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR) up

migrate-down:
	@echo "Reverting database migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR) down

migrate-up-test:
	@echo "Applying database test migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR_TEST) up

migrate-down-test:
	@echo "Reverting database test migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR_TEST) down

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

.PHONY: clean
