package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("We will learn how to handle get request today !! Welcome Boss :)")
	url := "http://localhost:3000/get"
	handleGetRequest(url)
}

func handleGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//handling response with 2 types

	//1 : usign ioutil package and string method
	/* fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength) */
	content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))

	//2: using strings package
	var responseContent strings.Builder
	responseContent.Write(content)
	fmt.Println(responseContent.String())
}
