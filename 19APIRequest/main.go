package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("We will learn how to handle get request today !! Welcome Boss :)")
	url := "http://localhost:3000/get"
	handleGetRequest(url)

	urlPost := "http://localhost:3000/post"
	handlePostReq(urlPost)

	urlPostForm := "http://localhost:3000/postForm"
	handlePostForm(urlPostForm)
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

func handlePostReq(url string) {
	data := strings.NewReader(`
	{
		"course" : "golang by mahesh"
	}
	`)

	response, err := http.Post(url, "application/json", data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	var responseContent strings.Builder
	responseContent.Write(content)
	fmt.Println("Repsonse : ", responseContent.String())

}

func handlePostForm(urlPostForm string) {
	data := url.Values{}
	data.Add("Name", "Mahesh")
	data.Add("Present", "true")
	fmt.Println("data content : ", data)
	response, err := http.PostForm(urlPostForm, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	var responseContent strings.Builder
	responseContent.Write(content)
	fmt.Println("response content : ", responseContent.String())
}
