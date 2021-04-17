package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(" This is Math Example ")
	intSum := 67 + 78 + 55
	fmt.Printf(" The of 67 + 78 + 55  is %d ", intSum)
	floatSum := 67.25 + 78.23 + 88.59
	fmt.Printf("\n The of 67.25 + 78.23 + 88.49  is %f ", floatSum)
	// Solving the Precision Problem
	fmt.Print("\n After Using Math.Round ", math.Round(floatSum))
	//Rounding to 2 Decimal Places
	fmt.Print("\n Rounding to 2 Decimal Places ", math.Round(floatSum*100)/100)
}
