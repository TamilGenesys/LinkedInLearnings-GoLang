package main

import (
	"fmt"
)

func main() {

	//New function will throw an error
	// object2 := new(map[string]int)
	// object2["hello"] = 12
	// fmt.Println(object2)

	object1 := make(map[string]int)
	object1["hello"] = 12
	fmt.Println(object1["hello"])
}
