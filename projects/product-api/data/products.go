package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
)

// swagger:model

type Product struct {

	// the id of this user
	// required: true
	// min: 1
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name" validate:"required"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price" validate:"required,gte=0"` // gte=0 means greater than or equal to 0
	SKU         string  `json:"sku" db:"sku" validate:"required,sku"`       // sku is a custom validation tag
	CreatedAt   string  `json:"-" db:"created_at"`                          // `json:"-"` means this field will not be included in the JSON output
	UpdatedAt   string  `json:"-" db:"updated_at"`
	DeletedAt   *string `json:"-" db:"deleted_at"` // Pointer to handle NULL values
}

// func to validate the Product struct
func (p *Product) ValidateProduct() error {
	// Create a new validator instance
	validate := validator.New()

	//register a custom validation function for the SKU field
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}
	return validate.Struct(p)
}

// custom validation function for SKU field
// use validator.FieldLevel interface to access the field being validated
func validateSKU(fl validator.FieldLevel) bool {
	//sku is of the form SKU-1234
	re := regexp.MustCompile("^SKU-[0-9]+$")
	return re.MatchString(fl.Field().String())
}

type Products []*Product

// ToJSON encodes the Products slice to JSON and writes it to the provided io.Writer
// returns an error if the encoding fails
// passing Products as a receiver so that we can call this method on Products type
func (p *Products) ToJSON(w io.Writer) error {
	// new encoder that will write to the provided io.Writer
	e := json.NewEncoder(w)
	//convert the Products slice to JSON and write it to the provided io.Writer
	return e.Encode(p)
}

// FromJSON decodes the JSON from the provided io.Reader into the Product struct
// Passing Product as a receiver so that we can call this method on Product type
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	//decode the JSON from the provided io.Reader into the Product struct
	return e.Decode(p)
}

// ProductRepository handles database operations for products
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Global repository instance (will be initialized in main.go)
var productRepo *ProductRepository

// InitializeRepository initializes the global repository
func InitializeRepository(db *sql.DB) {
	productRepo = NewProductRepository(db)
}

// GetProducts retrieves all products from the database
func GetProducts() (Products, error) {
	if productRepo == nil {
		return nil, fmt.Errorf("product repository not initialized")
	}

	query := `
		SELECT id, name, description, price, sku, created_at, updated_at, deleted_at 
		FROM products 
		WHERE deleted_at IS NULL 
		ORDER BY id`

	rows, err := productRepo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products Products
	for rows.Next() {
		var p Product
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

// AddProduct adds a new product to the database
func AddProduct(p *Product) error {
	if productRepo == nil {
		return fmt.Errorf("product repository not initialized")
	}

	query := `
		INSERT INTO products (name, description, price, sku) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at, updated_at`

	err := productRepo.db.QueryRow(query, p.Name, p.Description, p.Price, p.SKU).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	return nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(id int, p *Product) error {
	if productRepo == nil {
		return fmt.Errorf("product repository not initialized")
	}

	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, sku = $4, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $5 AND deleted_at IS NULL 
		RETURNING updated_at`

	err := productRepo.db.QueryRow(query, p.Name, p.Description, p.Price, p.SKU, id).Scan(&p.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return ErrProductNotFound
		}
		return fmt.Errorf("failed to update product: %w", err)
	}

	p.ID = id
	return nil
}

// FindProduct finds a product by ID
func FindProduct(id int) (*Product, error) {
	if productRepo == nil {
		return nil, fmt.Errorf("product repository not initialized")
	}

	query := `
		SELECT id, name, description, price, sku, created_at, updated_at, deleted_at 
		FROM products 
		WHERE id = $1 AND deleted_at IS NULL`

	var p Product
	err := productRepo.db.QueryRow(query, id).Scan(
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

// DeleteProduct soft deletes a product (sets deleted_at timestamp)
func DeleteProduct(id int) error {
	if productRepo == nil {
		return fmt.Errorf("product repository not initialized")
	}

	query := `
		UPDATE products 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := productRepo.db.Exec(query, id)
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

var ErrProductNotFound = fmt.Errorf("Product not found")
