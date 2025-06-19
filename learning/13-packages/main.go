//if package is same then we can share global vars and functions
package main

//run both files for common things to work-> go run packagefile1.go packagefile2.go

//this global var will be used in other file also inside packagefile2.go, cant declare infer type outside function
var globalVar = 10

func main() {
	commonFunc()
}