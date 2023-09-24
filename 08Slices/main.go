package main

import "fmt"

func main() {
	fmt.Println("welcome to the slices learnigs")
	fruits := []string{}
	fruits = append(fruits, "apple")
	fmt.Println("fruits : ", fruits)
	fmt.Println("length : ", len(fruits))
	fmt.Println("capacity : ", cap(fruits))
	fruits = append(fruits, "banana", "cherry", "orange", "dragon fruit", "mango")
	fmt.Println("fruits : ", fruits)
	fmt.Println("length : ", len(fruits))
	fmt.Println("capacity : ", cap(fruits))
	fmt.Println("slice of slices : ", fruits[:len(fruits)])

	//delete the element
	index := 2
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println("fruits : ", fruits)

}
