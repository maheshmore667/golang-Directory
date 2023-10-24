package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websites := []string{
		"https://google.com",
		"https://w3Schools.com",
		"https://www.tradingview.com",
	}
	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()

}

func getStatusCode(webiste string) {
	defer wg.Done()
	res, err := http.Get(webiste)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success!!", res.StatusCode)
	}
}
