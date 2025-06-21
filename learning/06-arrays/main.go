package main

import (
	"fmt"
	"sort"
)

func main() {
	var ages [4]int = [4]int{51,42,1,350}
	fmt.Println(ages[2])

	//short
	var ages2 = [10]int{1,2,340}
	fmt.Println(len(ages2))

	//infer type
	names := [3]string{"barry", "peter", "bruce"}
	fmt.Println(names, names[2])


	//slices (use array under the hood)-> we dont give the size so it can change, intially it contains 3 names, later we add 1
	namesSlices := []string{"barry", "peter", "bruce"}

	//append returns new array and we assign it to slice
	namesSlices = append(namesSlices, "tony")

	fmt.Println(namesSlices, namesSlices[3])
	

	// arr[x:y ] get slice from x to y
	nameRange := namesSlices[1:3]

	nameRange2 := namesSlices[2:]

	fmt.Println(nameRange, nameRange2)

	//sort, expects a slice and not an array, sort modifies the slice in place
	ages3 := []int{31,1,30,4}
	sort.Ints(ages3)
	fmt.Println(ages3)

	// The slicing operation ages[:] creates a slice that refers to the entire array. 
	//The slice shares the same underlying memory as the array, so changes made to the slice will be reflected in the array and vice versa.
	
	sort.Ints(ages[:])

	fmt.Println(ages)


	
}