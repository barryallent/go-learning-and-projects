package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsHandler struct {
	l *log.Logger
}

// NewProductsHandler This is like a constructor in java, it initializes the struct
func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l: l}
}

// GetProducts p is a receiver, it is like this in java
// we need to add a receiver to the method so that it can be called on the ProductsHandler struct,
// otherwise we won't be able to call it on the ProductsHandler instance
func (p *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	// get all products from the data package
	Products, err := data.GetProducts()
	if err != nil {
		http.Error(w, "Unable to retrieve products", http.StatusInternalServerError)
		return
	}

	// call the ToJSON method on Products to convert it to JSON
	err = Products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *ProductsHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// Get the product from the context, which was set by the MiddlewareProductValidation
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	// add the product to the data store
	err := data.AddProduct(product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to add product, err:", err), http.StatusInternalServerError)
		return
	}
}

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
		http.Error(w, "Unable to update product", http.StatusInternalServerError)
		return
	}
}

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
