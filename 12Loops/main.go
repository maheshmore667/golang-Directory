package main

import "fmt"

func main() {
	fmt.Println("Loops and it's structure")

	weekDays := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	// conventional for loop
	/* for i := 0; i < len(weekDays); i++ {
		fmt.Println(weekDays[i])
	} */

	//other way of using for loop

	for index, value := range weekDays {
		//continue keyword use
		/* if index == 3 {
			continue
		} */
		if index == 3 {
			//goto jumpedLabel
			break
		}
		fmt.Printf("index %v, value %v  \n", index, value)
	}

	//label
jumpedLabel:
	fmt.Println("Jumped to label jumpedLabel")
}
