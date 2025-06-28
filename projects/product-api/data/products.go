package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"` //this is to show as id in JSON not as ID
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"` // `json:"-"` means this field will not be included in the JSON output
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
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

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for pos, p := range productList {
		if p.ID == id {
			return p, pos, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
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
