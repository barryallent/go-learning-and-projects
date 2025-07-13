package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// ProductsHandler handles product-related HTTP requests
type ProductsHandler struct {
	l *log.Logger
}

// NewProductsHandler This is like a constructor in java, it initializes the struct
func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l: l}
}

// swagger:route GET / products listProducts
// Gets all products from the database
// responses:
//	200: productsResponse
//  500: errorResponse

// GetProducts returns all products
func (p *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	// get all products from the data package
	Products, err := data.GetProducts()
	if err != nil {
		http.Error(w, "Unable to retrieve products", http.StatusInternalServerError)
		return
	}

	// Set proper Content-Type header for JSON response
	w.Header().Set("Content-Type", "application/json")

	// call the ToJSON method on Products to convert it to JSON
	err = Products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route POST /product products createProduct
// Creates a new product
// responses:
//	201: productResponse
//  400: errorResponse
//  409: errorResponse
//  500: errorResponse

// AddProduct adds a new product to the database
// passing ProductsHandler as receiver so that we can call this method on ProductsHandler type
func (p *ProductsHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// Get the product from the context, which was set by the MiddlewareProductValidation
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	// add the product to the data store
	err := data.AddProduct(product)
	if err != nil {
		// Check for duplicate SKU error
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, fmt.Sprintf("Product with SKU '%s' already exists", product.SKU), http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("Unable to add product: %v", err), http.StatusInternalServerError)
		return
	}

	// Set proper Content-Type header and status for JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return the created product as JSON using json.NewEncoder
	encoder := json.NewEncoder(w)
	err = encoder.Encode(product)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route PUT /product/{id} products updateProduct
// Updates a product
// responses:
//	200: productResponse
//  400: errorResponse
//  404: errorResponse
//  409: errorResponse
//  500: errorResponse

// UpdateProducts updates an existing product
func (p *ProductsHandler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	// Get the product ID from the URL parameters using gorilla/mux
	vars := mux.Vars(r)

	// Convert the ID from string to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id to int", http.StatusBadRequest)
		return
	}

	// Get the product from the context, which was set by the MiddlewareProductValidation
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	// Call the UpdateProduct method from the data package to update the product
	err = data.UpdateProduct(id, product)

	// Check if the product was not found or if there was another error
	if err == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		// Check for duplicate SKU error
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, fmt.Sprintf("Product with SKU '%s' already exists", product.SKU), http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("Unable to update product: %v", err), http.StatusInternalServerError)
		return
	}

	// Set proper Content-Type header for JSON response
	w.Header().Set("Content-Type", "application/json")

	// Return the updated product as JSON
	encoder := json.NewEncoder(w)
	err = encoder.Encode(product)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// KeyProduct is used as a key for storing products in request context
type KeyProduct struct{}

// MiddlewareProductValidation to validate the product data before processing the request and passing it to the next handler
func (p *ProductsHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Println("Product validation middleware")

		//like doing this in java, Product p = new Product()
		product := &data.Product{}

		//get the product from the request body
		err := product.FromJSON(r.Body)

		if err != nil {
			http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// Validate the product struct using the ValidateProduct method that we defined in the data package
		err = product.ValidateProduct()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error Validating product: %s", err), http.StatusBadRequest)
			return
		}

		// Create a new context with the product and add it to the request so it can be accessed by the next handler
		//passing pointer to product
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)

		// Update the request with the new context
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// swagger:parameters createProduct updateProduct
type productParamsWrapper struct {
	// Product data
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters updateProduct
type productIDParamsWrapper struct {
	// Product ID
	// in: path
	// required: true
	ID int `json:"id"`
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products
	// in: body
	Body []data.Product
}

// A single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Product data
	// in: body
	Body data.Product
}

// Error response
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Error message
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}
