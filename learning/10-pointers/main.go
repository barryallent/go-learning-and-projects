package main

import "fmt"

func main() {
	name := "barry"
	m := &name

	fmt.Println("memory address of name is", m)
	//dereference pointer (get the value) using *
	fmt.Println("value at address m is", *m)

	fmt.Println("address at address m is", &*m)

	fmt.Println("value at address m is", *&*m)

	//change value at that address, so value of name will changed
	*m = "barry changed"
	fmt.Println(name)
}
