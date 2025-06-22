package data

import (
	"io"
	"time"
	"encoding/json"
)

type Product struct {
	ID          int        `json:"id"`    //this is to show as id in JSON not as ID
	Name        string     `json:"name"`
	Description string     `json:"description"` 
	Price       float64    `json:"price"`
	SKU		    string     `json:"sku"`
	CreatedAt   string     `json:"-"`       // `json:"-"` means this field will not be included in the JSON output
	UpdatedAt   string     `json:"-"`
	DeletedAt   string     `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

func GetProducts() Products {
	return ProductsList
}	

var ProductsList = []*Product {
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "frothy coffee with steamed milk",
		Price:       2.5,
		SKU:         "SKU001",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "strong coffee shot",
		Price:       1.5,
		SKU:         "SKU002", 
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	}, 
}