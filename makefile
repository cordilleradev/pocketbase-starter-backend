# Configuration
APP_NAME := pocketbase
BUILD_DIR := ./build
MAIN_FILE := ./main.go
DATA_DIR_DEV := ./pb_data
DATA_DIR_PROD := /var/data
PORT := 8080
# Use $PORT env variable in production (set by Render)

# Go build flags
GOFLAGS := -ldflags="-s -w"

# Ensure required directories exist
$(shell mkdir -p $(BUILD_DIR) $(DATA_DIR_DEV))

.PHONY: all build clean dev prod setup-prod backup help render-build render-start

# Default target
all: build

# Build the PocketBase application
build:
	@echo "Building $(APP_NAME)..."
	go build $(GOFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "Build completed successfully!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	@echo "Clean completed successfully!"

# Run PocketBase in development mode
dev:
	@echo "Starting $(APP_NAME) in development mode..."
	go run $(MAIN_FILE) serve --dev --dir=$(DATA_DIR_DEV) --http=0.0.0.0:$(PORT)

# For Render build step
render-build:
	@echo "Building for Render deployment..."
	go build $(GOFLAGS) -o app $(MAIN_FILE)

# For Render start command
render-start:
	@echo "Starting PocketBase on Render..."
	./app serve --dir=$(DATA_DIR_PROD) --http=0.0.0.0:$$PORT

# Create a backup of the production data
backup:
	@echo "Creating backup of production data..."
	TIMESTAMP=$$(date +%Y-%m-%d_%H-%M-%S); \
	tar -czf pocketbase_backup_$$TIMESTAMP.tar.gz $(DATA_DIR_PROD) && \
	echo "Backup created: pocketbase_backup_$$TIMESTAMP.tar.gz"

# Display help information
help:
	@echo "PocketBase Makefile Help"
	@echo "-----------------------"
	@echo "Available targets:"
	@echo "  all          : Same as 'build'"
	@echo "  build        : Compile the PocketBase application"
	@echo "  clean        : Remove build artifacts"
	@echo "  dev          : Run PocketBase in development mode (port $(PORT))"
	@echo "  setup-prod   : Set up production environment"
	@echo "  prod         : Run PocketBase in production mode (port $(PORT))"
	@echo "  render-build : Build for Render deployment"
	@echo "  render-start : Start command for Render deployment"
	@echo "  backup       : Create a backup of the production data"
	@echo "  help         : Display this help message"
