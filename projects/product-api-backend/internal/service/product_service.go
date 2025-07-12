package service

import (
	"fmt"
	"strings"

	"product-api/internal/models"
	"product-api/internal/repository"
)

// ProductService defines the interface for product business logic
type ProductService interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductByID(id int64) (*models.Product, error)
	CreateProduct(req *models.ProductCreateRequest) (*models.Product, error)
	UpdateProduct(id int64, req *models.ProductUpdateRequest) (*models.Product, error)
	DeleteProduct(id int64) error
}

// ProductService handles business logic for products
// This is the service layer in enterprise architecture - it sits between handlers and repository
// Like a service class in Java Spring framework
type productService struct {
	productRepo repository.ProductRepository
}

// NewProductService creates a new product service instance
// This is dependency injection - we inject the repository dependency
func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

// GetAllProducts retrieves all products
// This method delegates to the repository layer for data access
func (s *productService) GetAllProducts() ([]*models.Product, error) {
	products, err := s.productRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

// GetProductByID retrieves a product by ID
// This method adds business logic validation before calling the repository
func (s *productService) GetProductByID(id int64) (*models.Product, error) {
	// Business logic validation - check if ID is valid
	if id <= 0 {
		return nil, ErrInvalidProductID
	}

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		// Convert repository error to service error
		if err == repository.ErrProductNotFound {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}

	return product, nil
}

// CreateProduct creates a new product
// This method contains business logic for product creation
func (s *productService) CreateProduct(req *models.ProductCreateRequest) (*models.Product, error) {
	// Validate request using domain validation
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Business logic - Check if SKU already exists (business rule)
	existingProduct, err := s.productRepo.GetBySKU(req.SKU)
	if err != nil && err != repository.ErrProductNotFound {
		return nil, fmt.Errorf("failed to check SKU uniqueness: %w", err)
	}
	if existingProduct != nil {
		return nil, ErrSKUAlreadyExists
	}

	// Convert request to domain product
	product := req.ToProduct()

	// Create product in repository
	err = s.productRepo.Create(product)
	if err != nil {
		// Check for database-level duplicate constraint
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, ErrSKUAlreadyExists
		}
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

// UpdateProduct updates an existing product
// This method contains business logic for product updates
func (s *productService) UpdateProduct(id int64, req *models.ProductUpdateRequest) (*models.Product, error) {
	// Business logic validation - check if ID is valid
	if id <= 0 {
		return nil, ErrInvalidProductID
	}

	// Validate request using domain validation
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Business logic - Check if product exists
	existingProduct, err := s.productRepo.GetByID(id)
	if err != nil {
		if err == repository.ErrProductNotFound {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to get product for update: %w", err)
	}

	// Business logic - Check if SKU already exists for a different product
	if req.SKU != existingProduct.SKU {
		productWithSameSKU, err := s.productRepo.GetBySKU(req.SKU)
		if err != nil && err != repository.ErrProductNotFound {
			return nil, fmt.Errorf("failed to check SKU uniqueness: %w", err)
		}
		if productWithSameSKU != nil && productWithSameSKU.ID != id {
			return nil, ErrSKUAlreadyExists
		}
	}

	// Update product with new values using domain method
	existingProduct.UpdateProduct(req)

	// Persist changes to repository
	err = s.productRepo.Update(id, existingProduct)
	if err != nil {
		if err == repository.ErrProductNotFound {
			return nil, ErrProductNotFound
		}
		// Check for database-level duplicate constraint
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, ErrSKUAlreadyExists
		}
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return existingProduct, nil
}

// DeleteProduct soft deletes a product
// This method contains business logic for product deletion
func (s *productService) DeleteProduct(id int64) error {
	// Business logic validation - check if ID is valid
	if id <= 0 {
		return ErrInvalidProductID
	}

	// Delete product from repository
	err := s.productRepo.Delete(id)
	if err != nil {
		if err == repository.ErrProductNotFound {
			return ErrProductNotFound
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// Service level errors - these are business logic errors
var (
	ErrProductNotFound  = fmt.Errorf("product not found")
	ErrInvalidProductID = fmt.Errorf("invalid product ID")
	ErrSKUAlreadyExists = fmt.Errorf("product with this SKU already exists")
)
