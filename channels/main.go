package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel
	ch := make(chan int, 1)

	wg := &sync.WaitGroup{}

	// Add all the go routines to the wait group
	wg.Add(2)

	// go routines
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// chan<- is a send only channel
		ch <- 42
		ch <- 43
		wg.Done()
	}(ch, wg)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		// <-chan is a receive only channel
		value, ok := <-ch
		fmt.Println(value, ok)
		value, ok = <-ch
		fmt.Println(value, ok)
		wg.Done()
	}(ch, wg)

	// Wait for all the go routines to finish
	wg.Wait()

	// Close the channel
	close(ch)
}
