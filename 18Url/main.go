package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseName=angular&paid=true"

func main() {
	fmt.Println("Welcome to the url handlling studies")

	//parsing the url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	//reading diff values from the parsed url
	fmt.Println("url parts : ", result.Scheme)
	fmt.Println("url parts : ", result.Host)
	fmt.Println("url parts : ", result.Port())
	fmt.Println("url parts : ", result.Path)
	fmt.Println("url parts : ", result.RawQuery)

	//reading only queryparams from the parsed url
	qparams := result.Query()
	fmt.Println("qparams : ", qparams["courseName"])

	//constructing the url
	constructedUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		RawPath: "courseName=Reactjs",
	}

	fmt.Println("constructed url : ", constructedUrl.String())

}
