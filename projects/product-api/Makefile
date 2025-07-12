# Product API Makefile
.PHONY: help swagger run clean client clean-client

help: ## Show available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-12s %s\n", $$1, $$2}'

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

client: swagger ## Generate swagger client
	@echo "🔧 Generating Swagger client..."
	@if [ ! -f "$$(go env GOPATH)/bin/swagger" ]; then \
		echo "Installing swagger tool..."; \
		go install github.com/go-swagger/go-swagger/cmd/swagger@latest; \
	fi
	@if [ ! -d "../client-product-api" ]; then mkdir -p ../client-product-api; fi
	@cd ../client-product-api && \
		if [ ! -f "go.mod" ]; then \
			go mod init client-product-api; \
		fi && \
		$$(go env GOPATH)/bin/swagger generate client -f ../product-api/swagger.yaml -A product-api && \
		echo "📦 Installing dependencies..." && \
		go mod tidy
	@echo "✅ Client generated in ../client-product-api/ directory!"
	@echo "📁 Import path: client-product-api/client"
	@echo "🎯 Client is ready to use!"

clean-client: ## Remove generated client files
	@echo "🧹 Cleaning up client files..."
	rm -rf ../client-product-api/
	@echo "✅ Client cleanup complete!"

clean: ## Remove generated files
	@echo "🧹 Cleaning up..."
	rm -f swagger.yaml
	rm -rf ../client-product-api/
	@echo "✅ Cleanup complete!" 