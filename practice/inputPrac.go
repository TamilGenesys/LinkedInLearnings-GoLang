package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" Simple Type Conversion from String to Int ")
	strReader := bufio.NewReader(os.Stdin)
	fmt.Print(" Enter a number ")
	strValue, _ := strReader.ReadString('\n')
	fmt.Print(" String value is ", strValue)
	intValue, err := strconv.ParseInt(strings.TrimSpace(strValue), 10, 64)

	if err == nil {
		fmt.Print("Success Retrived ", intValue)
	} else {
		fmt.Print(" Failure ", err)
	}

}
