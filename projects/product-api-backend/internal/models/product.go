package models

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// swagger:model

// Product represents a product entity in our domain
type Product struct {

	// the id of this user
	// required: true
	// min: 1
	ID          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name" validate:"required"`
	Description string     `json:"description" db:"description"`
	Price       float64    `json:"price" db:"price" validate:"required,gte=0"` // gte=0 means greater than or equal to 0
	SKU         string     `json:"sku" db:"sku" validate:"required,sku"`       // sku is a custom validation tag
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`                 // `json:"-"` means this field will not be included in the JSON output
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Pointer to handle NULL values
}

// Products is a slice of Product pointers
type Products []*Product

// ProductCreateRequest represents the request for creating a product
type ProductCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gte=0"` // gte=0 means greater than or equal to 0
	SKU         string  `json:"sku" validate:"required,sku"`     // sku is a custom validation tag
}

// ProductUpdateRequest represents the request for updating a product
type ProductUpdateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gte=0"` // gte=0 means greater than or equal to 0
	SKU         string  `json:"sku" validate:"required,sku"`     // sku is a custom validation tag
}

// ToProduct converts ProductCreateRequest to Product
func (pcr *ProductCreateRequest) ToProduct() *Product {
	return &Product{
		Name:        pcr.Name,
		Description: pcr.Description,
		Price:       pcr.Price,
		SKU:         pcr.SKU,
	}
}

// UpdateProduct updates product with new values
func (p *Product) UpdateProduct(req *ProductUpdateRequest) {
	p.Name = req.Name
	p.Description = req.Description
	p.Price = req.Price
	p.SKU = req.SKU
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

// func to validate the ProductCreateRequest struct
func (pcr *ProductCreateRequest) Validate() error {
	// Create a new validator instance
	validate := validator.New()

	//register a custom validation function for the SKU field
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}
	return validate.Struct(pcr)
}

// func to validate the ProductUpdateRequest struct
func (pur *ProductUpdateRequest) Validate() error {
	// Create a new validator instance
	validate := validator.New()

	//register a custom validation function for the SKU field
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}
	return validate.Struct(pur)
}

// custom validation function for SKU field
// use validator.FieldLevel interface to access the field being validated
func validateSKU(fl validator.FieldLevel) bool {
	//sku is of the form SKU-1234
	re := regexp.MustCompile("^SKU-[0-9]+$")
	return re.MatchString(fl.Field().String())
}

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

// IsDeleted checks if the product is soft deleted
func (p *Product) IsDeleted() bool {
	return p.DeletedAt != nil
}

// MarkAsDeleted marks the product as deleted
func (p *Product) MarkAsDeleted() {
	now := time.Now()
	p.DeletedAt = &now
}
