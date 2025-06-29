# Product API Makefile
.PHONY: help swagger run clean

help: ## Show available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-10s %s\n", $$1, $$2}'

swagger: ## Generate Swagger documentation
	@echo "Generating Swagger documentation..."
	@if [ ! -f "$$(go env GOPATH)/bin/swagger" ]; then \
		echo "Installing swagger tool..."; \
		go install github.com/go-swagger/go-swagger/cmd/swagger@latest; \
	fi
	$$(go env GOPATH)/bin/swagger generate spec -o ./swagger.yaml --scan-models
	@echo "✅ Swagger documentation generated!"
	@echo "🌐 Access docs at: http://localhost:9080/docs"

run: ## Run the application
	@echo "Starting Product API server..."
	go run main.go

dev: swagger run ## Generate swagger docs and run application

clean: ## Remove generated files
	@echo "Cleaning up..."
	rm -f swagger.yaml
	@echo "✅ Cleanup complete!" 