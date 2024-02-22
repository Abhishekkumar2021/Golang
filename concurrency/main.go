package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex

var counter int = 0

func incrementCounter() {
	fmt.Println("Incrementing counter")
	// This region is a critical section so we need to lock it
	mutex.Lock()
	counter++
	mutex.Unlock()
	waitGroup.Done()
}

func decrementCounter() {
	fmt.Println("Decrementing counter")
	// This region is a critical section so we need to lock it
	mutex.Lock()
	counter--
	mutex.Unlock()
	waitGroup.Done()
}

func printString(s string) {
	for _, r := range s {
		fmt.Printf("%c ", r)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func getStatusCode(endpoint string) {
	defer waitGroup.Done()
	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Error in fetching the URL: %s\n", endpoint)
	}
	fmt.Printf("Status code for %s: %d\n", endpoint, resp.StatusCode)
}

func main() {
	go printString("ABCDEF")
	printString("123456")
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
	}

	for _, site := range websites {
		go getStatusCode(site)
		waitGroup.Add(1)
	}

	waitGroup.Add(2)
	go decrementCounter()
	go incrementCounter()

	// It's important to call waitGroup.Done() in the goroutine
	waitGroup.Wait()

	fmt.Println("Counter:", counter)
}
