package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {


	//need this because fmt.scan doesn't deal with spaces
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the input string \n")

	inputString, _ := consoleReader.ReadString('\n')
	inputString2 := strings.ToLower(inputString)

	//need to find the last character apart from the line feed
	lastChar  := inputString2[len(inputString2)-2:len(inputString2)-1]

	if strings.HasPrefix(inputString2, "i")  && strings.Contains(inputString2, "a") && lastChar == "n" {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}


