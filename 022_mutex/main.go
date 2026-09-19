package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex // creating a mutex

func increment() {
	mu.Lock()
	counter++
	defer mu.Unlock()
}

func main() {
	var wg sync.WaitGroup // creating a waitgroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		increment()
	}()

	go func() {
		defer wg.Done()
		increment()
	}()

	wg.Wait()
	fmt.Println(counter)
}

// Mutex -> mutual exclusion

/*
* WaitGroup -> wait for the goroutines to finish
* Channel -> Send this data to another goroutine
* Mutex -> Only one goroutine at a time
 */
