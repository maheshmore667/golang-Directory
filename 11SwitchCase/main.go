package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch case based dice game")
	rand.Seed(time.Now().UnixNano())
	randoNumber := rand.Intn(6) + 1

	switch randoNumber {
	case 1:
		fmt.Println("Enable to move further")
	case 2:
		fmt.Println("Move 2 spots ahead")
		//fallthrough : to run below test cases also
	case 3:
		fmt.Println("Move 3 spots ahead")
	case 4:
		fmt.Println("Move 4 spots ahead")
	case 5:
		fmt.Println("Move 5 spots ahead")
	case 6:
		fmt.Println("Move 6 spots ahead")

	}
}
