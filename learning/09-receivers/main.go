package main

import "fmt"

// Define the bill type
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// receive bill function, so can be only called from bill object like b.formatBill()
// this b bill passed as receiver will be copy of bill passed
// We can pass pointer (b *bill) to save memory so that if func is called multiple times we dont copy each time
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
	// Create a bill instance
	myBill := bill{
		name: "barry",
		items: map[string]float64{
			"item1": 12,
			"item2": 13,
			"item3": 5,
		},
		tip: 5.2,
	}

	// Use the receiver method
	fmt.Println(myBill.formatBill())
}
