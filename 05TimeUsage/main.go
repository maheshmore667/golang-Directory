package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time practises")
	currentTime := time.Now()
	fmt.Println("current time : ", currentTime)
	modifiedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("modified time : ", modifiedTime)

	createdDate := time.Date(2024, time.February, 12, 20, 00, 00, 00, time.Local)

	fmt.Println("created date is : ", createdDate.Format("01-02-2006 15:04:05 Monday"))

}
