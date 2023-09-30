package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling")
	content := "This tutorial is done by Mahesh. Please stay tuned to get more on this"

	//creating the file.
	file, err := os.Create("./fileLogs.txt")
	checkNilErr(err)

	//writing the file
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length of the file created is : ", length)

	defer file.Close()

	//reading the file
	dataBytes, err := ioutil.ReadFile("./fileLogs.txt")
	fmt.Println("File Data : ", string(dataBytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
