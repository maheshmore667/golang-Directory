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
	decodeJson()
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

func decodeJson() {

	decodedJson := []byte(`
	{
			"Name": "Mahesh",
			"Email": "mahesh@gmail.com",
			"Tags": ["web","js"]
	}`)
	isValid := json.Valid(decodedJson)
	var responseContent strings.Builder
	if isValid {
		fmt.Println("Json is valid")
		var newStudent Student
		json.Unmarshal(decodedJson, &newStudent)
		responseContent.Write(decodedJson)
		fmt.Println("new Student json : ", responseContent.String())

	} else {
		fmt.Println("Json is invalid")
	}

	//when you don't have struct defined

	var jsonDataFromWeb map[string]interface{}
	json.Unmarshal(decodedJson, &jsonDataFromWeb)
	responseContent.Write(decodedJson)
	fmt.Println("new Student json : ", responseContent.String())
}
