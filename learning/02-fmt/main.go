package main

import "fmt"

func main() {
	name := "barry"
	age := 24

	//print withour line
	fmt.Print("abc" + " ")
	fmt.Print("xyz \n")

	fmt.Println("My name is", name, "and my age is", age)

	//formatted string
	fmt.Printf("My name is %v and my age is %v \n", name, age)

	fmt.Printf("Age is of type %T \n", age)

	fmt.Printf("My score is %f \n", 29.5234)
	//rounded off to  2 decimal places only so use 0.2f
	fmt.Printf("My score is %0.2f \n", 29.5264)

	//SprintF (save formatted string)
	scoreString := fmt.Sprintf("My score is %0.2f \n", 29.5264)

	fmt.Println(scoreString)
}
