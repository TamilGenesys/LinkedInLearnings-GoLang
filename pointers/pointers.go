package main

import (
	"fmt"
)

func main() {
	text := "Am pointer1"
	pointer1 := &text
	fmt.Println(" The Address pointed by pointer1 is: ", pointer1)
	fmt.Println(" The value present in the addrss pointed by pointer1 is: ", *pointer1)
}
