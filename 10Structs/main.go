package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")
	Mahesh := Student{"Mahesh", "mahesh@test.mail", 25, true}

	fmt.Printf("Mahesh: %+v ", Mahesh)

}

type Student struct {
	Name    string
	Email   string
	Age     int
	Present bool
}
