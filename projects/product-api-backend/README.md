# 🚀 Product API with PostgreSQL and Swagger

A RESTful API for managing products, built with Go, PostgreSQL, and complete Swagger documentation.

## ✨ Features

- ✅ **Full CRUD Operations** - Create, Read, Update, Delete products
- ✅ **PostgreSQL Database** - Production-ready database integration
- ✅ **Swagger Documentation** - Interactive API documentation
- ✅ **Input Validation** - Request validation middleware
- ✅ **Error Handling** - Proper HTTP status codes and error messages
- ✅ **Soft Deletes** - Records marked as deleted, not removed
- ✅ **Auto Migration** - Database tables created automatically
- ✅ **Environment Config** - Flexible configuration via environment variables

## 🏗️ Project Structure

```
product-api/
├── config/                 # Configuration management
│   └── config.go
├── database/              # Database connection & migrations
│   └── connection.go
├── data/                  # Data models & repository
│   └── products.go
├── handlers/              # HTTP handlers with Swagger annotations
│   └── products.go
├── main.go               # Application entry point
├── swagger.yaml          # Generated Swagger specification
├── Makefile             # Build automation
├── go.mod               # Go module dependencies
└── README.md            # This file
```

## 🔧 Prerequisites

1. **Go** (version 1.19+)
2. **PostgreSQL** (version 12+)

## 🐘 PostgreSQL Setup

### 1. Install PostgreSQL

**macOS (using Homebrew):**
```bash
brew install postgresql
brew services start postgresql
```

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

### 2. Create Database and User

```bash
# Connect to PostgreSQL
psql postgres  # or: sudo -u postgres psql

# Run these SQL commands:
CREATE DATABASE product_api;
CREATE USER your_username WITH ENCRYPTED PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE product_api TO your_username;
\q
```

### 3. Configure Environment Variables

```bash
# Database Configuration
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=your_username
export DB_PASSWORD=your_password
export DB_NAME=product_api
export DB_SSL_MODE=disable

# Server Configuration
export SERVER_PORT=9080
```

## 🚀 Running the Application

### Quick Start
```bash
make dev      # Generate swagger docs and run the application
```

### Step by Step
1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Generate Swagger documentation:**
   ```bash
   make swagger
   ```

3. **Run the application:**
   ```bash
   make run
   ```

### Available Make Commands
```bash
make help     # Show all available commands
make swagger  # Generate Swagger documentation
make run      # Run the application
make dev      # Generate swagger docs and run application
make clean    # Remove generated files
```

The application will:
- Connect to PostgreSQL
- Create the `products` table automatically
- Insert sample data (Latte and Espresso)
- Start server on port 9080
- Serve Swagger documentation at `/docs`

## 📚 API Documentation

### 🌐 Interactive Documentation
- **Swagger UI:** [http://localhost:9080/docs](http://localhost:9080/docs)
- **Swagger Spec:** [http://localhost:9080/swagger.yaml](http://localhost:9080/swagger.yaml)

### 🛠️ API Endpoints

#### GET `/` - Get all products
```bash
curl http://localhost:9080/
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "Latte",
    "description": "frothy coffee with steamed milk",
    "price": 2.50,
    "sku": "SKU-001"
  }
]
```

#### POST `/product` - Create a new product
```bash
curl -X POST http://localhost:9080/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cappuccino",
    "description": "Coffee with steamed milk foam",
    "price": 3.50,
    "sku": "SKU-003"
  }'
```

#### PUT `/product/{id}` - Update a product
```bash
curl -X PUT http://localhost:9080/product/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Large Latte",
    "description": "Extra large latte with double shot",
    "price": 4.50,
    "sku": "SKU-001"
  }'
```

## 📊 Database Schema

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    sku VARCHAR(50) UNIQUE NOT NULL,  -- Stock Keeping Unit (must be unique)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL         -- For soft deletes
);
```

## 🔍 Input Validation

The API validates:
- **Required fields:** `name`, `price`, `sku`
- **Price validation:** Must be >= 0
- **SKU format:** Must match pattern `SKU-[0-9]+` (e.g., `SKU-001`)
- **Unique SKU:** Each product must have a unique Stock Keeping Unit

## 📝 Swagger Documentation

### Regenerating Documentation

After making changes to API annotations:

```bash
make swagger
```

### Adding New Endpoints

1. Add Swagger annotations to your handler:
```go
// swagger:route GET /products/{id} products getProduct
// Get a product by ID
// responses:
//   200: productResponse
//   404: errorResponse
func (p *ProductsHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    // implementation
}
```

2. Regenerate documentation:
```bash
make swagger
```

## 🐛 Troubleshooting

### Database Connection Issues
- Ensure PostgreSQL is running: `brew services list | grep postgresql`
- Test connection: `psql -d product_api -U your_username`
- Check environment variables

### Port Conflicts
- Change `SERVER_PORT` environment variable
- Default port is 9080

### Swagger Generation Issues
- Ensure swagger tool is installed: `go install github.com/go-swagger/go-swagger/cmd/swagger@latest`
- Check Go PATH: `echo $GOPATH/bin`
- Verify annotations syntax in handlers

### SKU Duplicate Errors
- SKUs must be unique across all products
- Use format: `SKU-001`, `SKU-002`, etc.
- Check existing products: `curl http://localhost:9080/`

## 🧪 Testing

### Test Product Creation
```bash
# Should succeed
curl -X POST http://localhost:9080/product \
  -H "Content-Type: application/json" \
  -d '{"name": "Mocha", "price": 4.0, "sku": "SKU-999"}'

# Should fail (duplicate SKU)
curl -X POST http://localhost:9080/product \
  -H "Content-Type: application/json" \
  -d '{"name": "Another Latte", "price": 3.0, "sku": "SKU-001"}'
```

### Test Validation
```bash
# Should fail (invalid SKU format)
curl -X POST http://localhost:9080/product \
  -H "Content-Type: application/json" \
  -d '{"name": "Test", "price": 2.0, "sku": "INVALID"}'
```

## 🎯 Development Tips

1. **Always regenerate Swagger docs** after changing API annotations
2. **Use meaningful SKUs** for easier inventory management
3. **Test error cases** to ensure proper error handling
4. **Check Swagger UI** to validate your API documentation
5. **Use environment variables** for configuration instead of hardcoding

## 🔄 Go Project Lifecycle

### 📋 **Development Phase**

```bash
# 1. Initialize new Go module
go mod init product-api-backend

# 2. Add dependencies 
go get github.com/gorilla/mux
go get github.com/lib/pq

# 3. Manage dependencies
go mod tidy          # Add missing, remove unused
go mod download      # Download modules to local cache
go mod verify        # Verify dependencies

# 4. Development workflow
go run main.go       # Run directly from source
go fmt ./...         # Format all Go files
go vet ./...         # Static analysis (find bugs)
```

### 🔨 **Build Phase**

```bash
# 1. Build executable
go build                    # Build in current directory
go build -o bin/app        # Build with custom name/path
go build ./...             # Build all packages

# 2. Cross-compilation
GOOS=linux GOARCH=amd64 go build -o bin/app-linux
GOOS=windows GOARCH=amd64 go build -o bin/app.exe
GOOS=darwin GOARCH=arm64 go build -o bin/app-mac

# 3. Production build (optimized)
go build -ldflags="-s -w" -o bin/app    # Strip debug info
CGO_ENABLED=0 go build                  # Static binary
```

### 🧪 **Testing Phase**

```bash
# 1. Run tests
go test ./...              # All packages
go test -v ./...           # Verbose output
go test -run TestName      # Specific test

# 2. Advanced testing
go test -race ./...        # Race condition detection
go test -cover ./...       # Code coverage
go test -bench=.           # Benchmark tests
go test -count=5 ./...     # Run tests multiple times

# 3. Coverage analysis
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out     # View in browser
```

### 📚 **Documentation Phase**

```bash
# 1. Generate Go docs
go doc                     # Show package docs
godoc -http=:6060         # Start documentation server

# 2. API Documentation (our project)
make swagger              # Generate Swagger docs
make client               # Generate API client

# 3. README and examples
# Keep README.md updated
# Add code examples
# Document environment variables
```


### 📊 **Project-Specific Lifecycle (Our Product API)**

```bash
# 1. Setup & Development
make help                 # Show available commands
export DB_USER=postgres   # Configure environment
export DB_PASSWORD=secret
export DB_NAME=product_api

# 2. Database Setup
make db-setup            # Show setup instructions
psql -c "CREATE DATABASE product_api;"

# 3. Development Workflow
make dev                 # Generate docs + run server
make swagger            # Update API documentation
make client             # Generate/update client

# 4. Testing API
curl http://localhost:9080/           # Test endpoints
curl http://localhost:9080/docs       # View documentation

# 5. Build & Deploy
make build              # (if added to Makefile)
make clean              # Clean generated files
make clean-client       # Clean client files

# 6. Client Integration
cd ../client-product-api-backend-backend
go run example.go       # Test generated client
```

### 🎯 **Best Practices**

1. **Version Control**
   ```bash
   git tag v1.0.0          # Tag releases
   git push --tags         # Push tags
   ```

2. **Environment Management**
   ```bash
   # Use .env files for local development
   # Environment variables for production
   # Never commit secrets
   ```

3. **Code Quality**
   ```bash
   # Always run before committing:
   go fmt ./...
   go vet ./...
   go test ./...
   golangci-lint run
   ```

4. **Documentation**
   ```bash
   # Keep README updated
   # Document API changes
   # Update Swagger annotations
   make swagger  # Regenerate docs
   ```

## 🚀 Next Steps

Consider adding:
- [ ] DELETE endpoint for products
- [ ] Product categories  
- [ ] Authentication & authorization
- [ ] Unit tests with table-driven tests
- [ ] Integration tests
- [ ] Docker containerization
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Rate limiting middleware
- [ ] Caching with Redis
- [ ] Monitoring with Prometheus
- [ ] Logging with structured logs

---

**Happy coding! 🎉** 