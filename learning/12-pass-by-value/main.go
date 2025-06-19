package main

import "fmt"

// Pass by value example
func modifyValue(x int) {
	x = x + 10
	fmt.Println("Inside modifyValue:", x)
}

// Pass by reference example (using pointer)
func modifyPointer(x *int) {
	*x = *x + 10
	fmt.Println("Inside modifyPointer:", *x)
}

func main() {
	fmt.Println("=== Pass by Value vs Pass by Reference ===")

	// Pass by value
	value := 5
	fmt.Println("Original value:", value)
	modifyValue(value)
	fmt.Println("After modifyValue:", value) // Value unchanged

	fmt.Println()

	// Pass by reference (pointer)
	pointerValue := 5
	fmt.Println("Original pointer value:", pointerValue)
	modifyPointer(&pointerValue)
	fmt.Println("After modifyPointer:", pointerValue) // Value changed

	fmt.Println()

	// Demonstrate with slices (pass by reference)
	slice := []int{1, 2, 3}
	fmt.Println("Original slice:", slice)
	modifySlice(slice)
	fmt.Println("After modifySlice:", slice) // Slice changed

	fmt.Println()

	// Demonstrate with arrays (pass by value)
	array := [3]int{1, 2, 3}
	fmt.Println("Original array:", array)
	modifyArray(array)
	fmt.Println("After modifyArray:", array) // Array unchanged
}

func modifySlice(s []int) {
	s[0] = 100
	fmt.Println("Inside modifySlice:", s)
}

func modifyArray(a [3]int) {
	a[0] = 100
	fmt.Println("Inside modifyArray:", a)
}
