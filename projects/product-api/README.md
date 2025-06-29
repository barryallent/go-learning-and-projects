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

## 🚀 Next Steps

Consider adding:
- [ ] DELETE endpoint for products
- [ ] Product categories
- [ ] Authentication & authorization
- [ ] Rate limiting
- [ ] Caching with Redis
- [ ] Unit tests
- [ ] Docker containerization
- [ ] CI/CD pipeline

---

**Happy coding! 🎉** 