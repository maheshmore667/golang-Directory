package main

import "fmt"

func main() {
	fmt.Println("welcome to the function studies")
	greet()
	sum := adder(23, 67)
	fmt.Println("sum : ", sum)

	multiSum, message := multiAdder(12, 3, 56, 78, 89)
	fmt.Println(message)
	fmt.Println("", multiSum)

}

//function with parameters
func adder(value1 int, value2 int) int {
	return value1 + value2
}

//function with multiple arguments
func multiAdder(values ...int) (int, string) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "sum from multiAdder : "

}

func greet() {
	fmt.Println("welcome to the greet")
}
