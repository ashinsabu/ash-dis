# Binary names
SERVER_BIN=bin/ashdis-server
CLI_BIN=bin/ashdis-cli

# Directories
SERVER_DIR=cmd/server
CLI_DIR=cmd/cli

# Build the server binary
build-server:
	@echo "🔨 Building Ashdis Server..."
	@go build -o $(SERVER_BIN) ./$(SERVER_DIR)

# Build the CLI binary
build-cli:
	@echo "🔨 Building Ashdis CLI..."
	@go build -o $(CLI_BIN) ./$(CLI_DIR)

# Build both server and CLI
build: build-server build-cli

# Run the server
run-server: build-server
	@echo "🚀 Starting ashdis Server..."
	@$(SERVER_BIN) --port=6370

# Clean up compiled binaries
clean:
	@echo "🧹 Cleaning up binaries..."
	@rm -f $(SERVER_BIN) $(CLI_BIN)

# Help menu
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build-server    Build the Ashdis server binary"
	@echo "  build-cli       Build the Ashdis CLI binary"
	@echo "  build           Build both server and CLI"
	@echo "  run-server      Build and run the server"
	@echo "  clean           Remove binaries"
	@echo "  help            Show this help message"
