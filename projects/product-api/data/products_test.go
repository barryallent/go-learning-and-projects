package data

import "testing"

func TestProductValidation(t *testing.T) {
	p := &Product{
		Name:  "Test Product",
		Price: 0,
		SKU:   "SKU-12",
	}

	err := p.ValidateProduct()

	if err != nil {
		t.Fatal(err)
	}

}
