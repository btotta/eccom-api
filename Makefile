# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."

	@go build -o main cmd/api/main.go

# Run the application
run:
	@echo "Initializing Swagger..."
	@swag init -g internal/server/routes.go -o docs
	@echo "Running..."
	@go run cmd/api/main.go

# Run test db and execute all tests in the tests folder
test:
	@clear
	@echo "Testing..."
	# Verifica se o banco de dados está rodando
	@if ! nc -z localhost 5400; then \
	    echo "Banco de dados não está rodando. Iniciando o banco de dados..."; \
	    docker-compose -f tests/docker-compose-db-test.yml up -d db; \
	    echo "Aguardando o banco de dados iniciar..."; \
	    sleep 10; \
	fi
	@echo "Executando testes..."
	@go test ./tests/... -v -count=1 -p 1

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main
	@rm -rf tmp

# Live Reload
watch:
	@if command -v air > /dev/null; then \
		swag init -g internal/server/routes.go -o docs; \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Install necessary tools and dependencies
install:
	@echo "Installing swag..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Installing air..."
	@go install github.com/cosmtrek/air@latest
	@echo "Installing dlv..."
	@go install -v github.com/go-delve/delve/cmd/dlv@latest
	@echo "Installing project dependencies..."
	@go mod tidy
	@echo "All necessary tools and dependencies have been installed."

.PHONY: all build run test clean watch install
