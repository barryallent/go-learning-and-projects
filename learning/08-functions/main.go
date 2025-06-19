package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greetings("Barry")

	names := []string{"barry", "tony", "bruce"}
	
	cycleNames(names, greetings)
	cycleNames(names, greetingsBye)

	area := areaCircle(4.2)

	fmt.Println("Area of circle is", area)

	name1, name2 := separateInitials("Barry Allen")

	fmt.Println(name1, name2)

	name11, name22 := separateInitials("Bruce")

	fmt.Println(name11, name22)
}


func greetings(name string) {
	fmt.Println("Hello", name)
}

func greetingsBye(name string) {
	fmt.Println("Bye", name)
} 


//passing func as argument to function, function expects a single string param so we can pass that type of function only
func cycleNames(names []string, f func(string)) {
	for i:=0; i < len(names); i++ {
		f(names[i])
	}
}


func areaCircle(r float32) float32 {
	return math.Pi * r * r 
}

//function with multiple return types
func separateInitials(s string) (string, string) {

	//convert a string to string array
	names := strings.Split(s, " ")

	//get the 1st char of both the strings

	if len(names) > 1 {
		return string(names[0][:1]), string([]rune(names[1])[:1])
	} else {
		return string(names[0][:1]), "_"
	}
	
	
}