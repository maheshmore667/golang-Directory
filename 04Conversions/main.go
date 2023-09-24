package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the conversions tutorial")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating to the tutorial : ")
	userInput, _ := reader.ReadString('\n')
	rating, _ := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
	fmt.Println("This is your rating afetr adding one : ", rating+1)
}
