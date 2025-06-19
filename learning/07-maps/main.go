package main

import "fmt"

func main() {

	//declare a map with key as type string and value as type float64
	menu := map[string]float64 {
		"item1" : 12,
		"item2" : 13,
		"item3" : 5,
	}

	//get the map and specific value
	fmt.Println(menu, menu["ite m1"])


	//iterate a map  
	for k, v := range menu {
		fmt.Printf("%v - %v \n", k,v)
	}
}