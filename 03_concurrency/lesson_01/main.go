package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printNumber() {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	wg.Add(5)

	go printNumber()
	go printNumber()
	go printNumber()
	go printNumber()
	go printNumber()

	alphabets := []string{"A", "B", "C", "D"}
	for _, char := range alphabets {
		fmt.Println(char)
	}

	wg.Wait()
}
