package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Date Time Example")
	currTime := time.Now()
	fmt.Println("Current Time is ", currTime)
	fmt.Println(currTime.Format(time.RFC1123))
}
