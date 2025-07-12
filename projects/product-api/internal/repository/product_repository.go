package repository

import (
	"database/sql"
	"fmt"

	"product-api/internal/models"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	GetAll() ([]*models.Product, error)
	GetByID(id int64) (*models.Product, error)
	Create(product *models.Product) error
	Update(id int64, product *models.Product) error
	Delete(id int64) error
	GetBySKU(sku string) (*models.Product, error)
}

// ProductRepository handles database operations for products
type productRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

// GetAll retrieves all products from the database
func (r *productRepository) GetAll() ([]*models.Product, error) {
	query := `
		SELECT id, name, description, price, sku, created_at, updated_at, deleted_at 
		FROM products 
		WHERE deleted_at IS NULL 
		ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.SKU,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

// GetByID retrieves a product by ID from the database
func (r *productRepository) GetByID(id int64) (*models.Product, error) {
	query := `
		SELECT id, name, description, price, sku, created_at, updated_at, deleted_at 
		FROM products 
		WHERE id = $1 AND deleted_at IS NULL`

	var p models.Product
	err := r.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.SKU,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to find product: %w", err)
	}

	return &p, nil
}

// Create adds a new product to the database
func (r *productRepository) Create(product *models.Product) error {
	query := `
		INSERT INTO products (name, description, price, sku) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.SKU).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	return nil
}

// Update updates an existing product in the database
func (r *productRepository) Update(id int64, product *models.Product) error {
	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, sku = $4, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $5 AND deleted_at IS NULL 
		RETURNING updated_at`

	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.SKU, id).Scan(&product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return ErrProductNotFound
		}
		return fmt.Errorf("failed to update product: %w", err)
	}

	product.ID = id
	return nil
}

// Delete soft deletes a product (sets deleted_at timestamp)
func (r *productRepository) Delete(id int64) error {
	query := `
		UPDATE products 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

// GetBySKU retrieves a product by SKU from the database
func (r *productRepository) GetBySKU(sku string) (*models.Product, error) {
	query := `
		SELECT id, name, description, price, sku, created_at, updated_at, deleted_at 
		FROM products 
		WHERE sku = $1 AND deleted_at IS NULL`

	var p models.Product
	err := r.db.QueryRow(query, sku).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.SKU,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to find product by SKU: %w", err)
	}

	return &p, nil
}

// Common errors
var (
	ErrProductNotFound = fmt.Errorf("product not found")
)
