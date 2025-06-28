package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-api/data"
	"strconv"
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
	Products := data.GetProducts()

	// call the ToJSON method on Products to convert it to JSON
	err := Products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *ProductsHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// data.Product is a struct that we defined in the data package
	product := &data.Product{}

	// get the product data from the request body and unmarshal it into the product struct
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	// add the product to the data store
	data.AddProduct(product)
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
	// Create a new Product instance to hold the data from the request body
	product := &data.Product{}

	// Use the FromJSON method to populate the product instance from the request body
	err = product.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	// Call the UpdateProduct method from the data package to update the product
	error := data.UpdateProduct(id, product)

	// Check if the product was not found or if there was another error
	if error == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(w, "Unable to update product", http.StatusInternalServerError)
		return
	}
}
