package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var signals = []string{}

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
		//trying to lock the memory block for this operation using mutex
		mutex.Lock()
		signals = append(signals, webiste)
		//trying to unlock the memory block which is reserved by mutex
		mutex.Unlock()
		fmt.Println("success!!", res.StatusCode)
		fmt.Println(signals)
	}
}
