package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")

	fmt.Println("Please enter anything on the keyboard : ")
	reader := bufio.NewReader(os.Stdin)
	//comma ok syntax | comma err syntax
	userInput, _ := reader.ReadString('\n')
	fmt.Println(userInput)

}
