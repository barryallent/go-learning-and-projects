package main

import "fmt"


func main() {
	age := 35
	fmt.Println(age>50)


	//else if formatting is mattering, cant go to new line
	//The else keyword must be on the same line as the closing curly brace (}) of the preceding block.
	if age < 35 {
		fmt.Println("age is less than 35")
	} else if age > 35 {
		fmt.Println("age is greater than 35")
	} else {
		fmt.Println("age is random")
	}
}