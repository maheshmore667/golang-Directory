package main

import (
	"fmt"
	"sync"
)

var score = []int{}

func main() {
	fmt.Println("Race Condition in Mutex and Wait Groups")
	wg := &sync.WaitGroup{}
	mt := &sync.Mutex{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {

		fmt.Println("first Func")
		mt.Lock()
		score = append(score, 1)
		mt.Unlock()
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("second Func")
		mt.Lock()
		score = append(score, 2)
		mt.Unlock()
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("third Func")
		mt.Lock()
		score = append(score, 3)
		mt.Unlock()
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
