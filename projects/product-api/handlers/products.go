package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-api/data"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	Products := data.GetProducts()

	// call the ToJSON method on Products to convert it to JSON
	err := Products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	data.AddProduct(product)
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id to int", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Products")

	product := &data.Product{}

	perror := product.FromJSON(r.Body)

	if perror != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	error := data.UpdateProduct(id, product)

	if error == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(w, "Unable to update product", http.StatusInternalServerError)
		return
	}
}
