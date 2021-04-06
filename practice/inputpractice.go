package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readerEntity := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name ")
	userName, _ := readerEntity.ReadString('\n')
	fmt.Printf("Hi %s, Welcome to SwissBank ", strings.TrimRight(userName, "\r\n"))

	fmt.Printf("Please Enter the Ac\\No ")
	userAccount, _ := readerEntity.ReadString('\n')
	userAccountNum, err := strconv.ParseInt(strings.TrimSpace(userAccount), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Your AC is verified %d", userAccountNum)
	}

}
