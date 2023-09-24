package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers playground")
	var ptr *int
	fmt.Println("value of the ptr : ", ptr)
	myNumber := 23
	fmt.Println("value of address of mynumber : ", &myNumber)
	ptr = &myNumber
	fmt.Println("value of the ptr : ", ptr)
	fmt.Println("*ptr : ", *ptr)

	*ptr = *ptr + 5
	fmt.Println("*ptr : ", *ptr)

}
