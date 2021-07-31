package main

import (
	"fmt"
	"myapp/packageone"
)

func main() {
	var myString = "Test"
	fmt.Println(myString)

	newString := packageone.PublicVar
	fmt.Println("From packageone", newString)

	packageone.Exported()

}
