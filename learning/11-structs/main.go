package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"item1": 12,
			"item2": 13,
			"item3": 5,
		},
		tip: 5.2,
	}
	return b

}

// Add the formatBill method
func (b bill) formatBill() string {
	fs := "Bill breakdown: \n"
	total := 0.0
	for k, v := range b.items {
		fs += fmt.Sprintf("%v : %v \n", k, v)
		total += v
	}
	fs += fmt.Sprintf("Total: %v", total)
	return fs
}

func main() {

	//create new bill
	myBill := newBill("barry")
	fmt.Println(myBill)

	//format bill
	fmt.Println(myBill.formatBill())
}
