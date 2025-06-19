package main

import "fmt"


//this is a common function which is also used in packagefile1.go
func commonFunc() {
	fmt.Println("hello")
	fmt.Println(globalVar)
}

