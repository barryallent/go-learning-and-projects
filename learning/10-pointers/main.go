package main
import "fmt"

func main() {
	name := "barry"
	fmt.Println("memory address is", &name)

	m := &name
	//dereference pointer (get the value) using *
	fmt.Println("value at address m is", *m)

	//change value at that address, so value of name will changed
	*m = "barry changed"
	fmt.Println(name)
}