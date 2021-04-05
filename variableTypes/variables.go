package main

import (
	"fmt"
)

func main() {
	//Explicit declaration of int
	//https://golang.org/pkg/fmt/ for more related to %<placeholder>

	var aInt int = 24
	fmt.Printf("The value of aInt is %d", aInt)

	// when var keyword is used, this variable can be used within function only
	bInt := 34
	fmt.Printf("\nThe value of bInt is %d", bInt)

	//Find the Type Print type
	aFloat := 24.848484
	fmt.Printf("\nThe type of %f is %T", aFloat, aFloat)

	abool := true
	fmt.Printf("\nThe type of %t is %T", abool, abool)

}
