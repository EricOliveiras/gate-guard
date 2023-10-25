#!/bin/bash

# Run Docker Compose
echo "Starting Docker Compose..."
docker-compose up -d

# Wait a few seconds to ensure the database is ready
sleep 5

# Apply database migrations
make migrate-up

# Compile the Go application
make build

# Run tests
make go-test

# Run the application
make run
