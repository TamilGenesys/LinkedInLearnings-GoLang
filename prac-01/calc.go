package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" This is a Calculator Program ")

	//Step 1 : Create the ReaderEntity
	readerEntity := bufio.NewReader(os.Stdin)

	//Step2 : Prompt the user to enter the value on int 1
	fmt.Print("Please enter the value of integer1 ")
	input1, _ := readerEntity.ReadString('\n')

	//Step 3 : Type convert the derived value to Integer
	integer1, err := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	//Step4 : Prompt the user to enter the value on int 1
	fmt.Print("Please enter the value of integer2 ")
	input2, _ := readerEntity.ReadString('\n')

	//Step 3 : Type convert the derived value to Integer
	integer2, err := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)

	if err != nil {
		fmt.Println(err)
	}
	sum := integer1 + integer2

	fmt.Printf("The Sum of %d and %d is %d", integer1, integer2, sum)
}
