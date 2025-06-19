package main

import "fmt"

func main() {
	//1st way to declare where we tell the type
	var create1 string = "var1"
	//2nd way to declare where we dont tell the type and go automatically understands
	var create2 = "var2"

	//defined the variable as string so later it will be a string only, its a empty string
	var create3 string
	fmt.Println(create1, create2, create3)

	create1 = "var1 changed"
	create3 = "var3"

	fmt.Println(create1, create3)

	//can declar via colon also instead of var. But used for new variable only. Cant use this notation outside of a function
	create4 := "new var initialized via colon"
	
	//will give error because create1 is already defined.
	//create1 := "new var"
	
	fmt.Println(create4)

	integersVars()

	floatVars()
	 
}

func integersVars() {
	var var1 int = 10
	var var2 = 20000
	var3 := 301
	fmt.Println(var1, var2, var3)

	//bits and memory
	//8 bit integer via int8, value from -128 to 127
	var var4 int8 = 127

	var var5 int16 = 30000

	fmt.Println(var4, var5)

	//unsigned int, unsignted int8 is from 0 to 255 because negatives are not allowed.
	var var6 uint8 = 255 
	fmt.Println(var6)

}

func floatVars() {
	var var1 float32 = 10.34
	var var2 float64 = 20921.21211213131
	
	//infer the type
	var3 := 9121.21212112
	fmt.Println(var1, var2, var3)
}