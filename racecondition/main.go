package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var score = []int{}
	// IIFE (Immediately Invoked Function Expression)
	wg.Add(3)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("First goroutine")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Second goroutine")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Third goroutine")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()

	fmt.Println(score)
}
