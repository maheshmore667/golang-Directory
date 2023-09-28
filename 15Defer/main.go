package main

import "fmt"

func main() {
	fmt.Println("Defer studies")
	defer fmt.Println("Hello")
	defer fmt.Println("Two")
	defer fmt.Println("One")

	fmt.Println("World")
	printNums()
}

func printNums() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//op :
/* Defer studies
world 4
3
2
1
0
one
two
Hello */
