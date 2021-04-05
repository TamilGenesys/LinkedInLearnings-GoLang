package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readerEntity := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name ")
	userName, _ := readerEntity.ReadString('\n')
	// fmt.Println() https://stackoverflow.com/questions/44448384/how-remove-n-from-lines
	userName = strings.TrimRight(userName, "\r\n")
	fmt.Printf("Hi %s, welcome to GeeksLand  ", userName)

}
