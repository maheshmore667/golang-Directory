package main

import "fmt"

func main() {
	Mahesh := User{Name: "Mahesh", Mail: "dev@test.com", Age: 25, Status: true}
	fmt.Println("User is : ", Mahesh)
	//calling method on the struct
	Mahesh.GetStatus()

	//manipulating the struct properties
	Mahesh.ChangeEmail()
}

type User struct {
	Name   string
	Mail   string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("User is active : ", u.Status)
}

func (u User) ChangeEmail() {
	u.Mail = "go@dev.com"
	fmt.Println("Mail of the user is : ", u.Mail)
}
