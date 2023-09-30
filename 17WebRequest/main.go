package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request studies")

	//invoking the web get request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// reading the response body got from server
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//parsing the content
	content := string(body)
	fmt.Println("resonse body is : ", content)
	defer response.Body.Close()
}
