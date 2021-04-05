package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Reader Entity, gets input from the standard input
	readerEntity := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the UserName ")
	//We are reading from the standardInput with delimter newline
	userName, _ := readerEntity.ReadString('\n')
	fmt.Printf("Hi %s, welcome to SBI ", strings.TrimRight(userName, "\r\n"))

	fmt.Print("\nPlease enter the Account Number ")
	userAccount, _ := readerEntity.ReadString('\n')
	accountNum, err := strconv.ParseInt(strings.TrimSpace(userAccount), 10, 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Thanks for confirming your AC\\Number ", accountNum)
	}

}
