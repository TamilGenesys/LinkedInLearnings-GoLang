package main

import (
	"fmt"
)

func main() {
	// Declaring an Array

	var array1 [5]int

	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println("The contents of Array1 is ", array1)

	//Single line declartion and assignment
	array2 := [3]string{"one", "two", "three"}
	fmt.Println("The contents of Array2 is ", array2)

}
