package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, World!")
}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func task(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the waitgroup counter when the goroutine finishes
	fmt.Println("Task finished")
}

func main() {
	// go sayHello()
	// fmt.Println("Main")

	// go count("A")
	// go count("B")

	var wg sync.WaitGroup // making a waitgroup variable
	wg.Add(1)             // adding 1 goroutine
	go task(&wg)          // start goroutine
	wg.Wait()             // wait for goroutine to finish (until the counter reaches 0)

	fmt.Println("Main finished")

	// time.Sleep(3 * time.Second)
}

// goroutine is a lightweight thread of execution that allows multiple functions to run concurrently.
