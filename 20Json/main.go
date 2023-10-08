package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Student struct {
	Name  string
	Email string
	Tags  []string
}

func main() {
	fmt.Println("Welcome to json studies")
	handleJson()
}

func handleJson() {
	bluePrintData := []Student{
		{"Mahesh", "mahesh@gmail.com", []string{"web", "js"}},
		{"Mahesh", "mahesh@gmail.com", []string{"web", "js"}},
	}

	fmt.Println("bluePrint : ", bluePrintData)
	jsonData, err := json.MarshalIndent(bluePrintData, "", "\t")

	if err != nil {
		panic(err)
	}
	var responseContent strings.Builder
	responseContent.Write(jsonData)
	fmt.Println("json : ", responseContent.String())

}
