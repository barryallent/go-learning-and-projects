package main

import (
	"fmt"
	"strings"
)

func main() {
	argo := "string is called news"

	//contains returns true or false 
	fmt.Println(strings.Contains(argo,"called"))

	//replace all returns new string
	argo2 := strings.ReplaceAll(argo, "is", "are")

	argo3 := strings.ToUpper(argo)

	fmt.Println(argo2, argo3)
}