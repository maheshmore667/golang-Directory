package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to array studies")
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "cherry"
	fruits[3] = "cherry"
	fruits[3] = "cherry"

	fmt.Println("Fruit list is : ", fruits)
	fmt.Println("length of the fruit list is : ", len(fruits))

	var sports = [4]string{"cricket", "khokho", "carrum"}

	fmt.Println("sport list is : ", sports)

}
