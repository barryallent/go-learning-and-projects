package main

import "fmt"

func main() {
	x := 0

	//same as a while loop but we use for in go
	for x<5 {
		fmt.Println(x)
		x++
	}

	//another way 
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is =",i)
	}


	//iterate in array
	names := []string {"name1", "name2", "name3"}

	for i := 0; i< len(names); i++ {
		fmt.Println(names[i])
	}

	//
	for index, value := range names {
		fmt.Printf("index is %v and value is %v \n", index, value)
	}
}