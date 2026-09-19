package main

import "fmt"

func main() {
	ch := make(chan int) // creating a channel of type int

	go func() {
		ch <- 10 // sending a value to channel
	}()

	value := <-ch // receiving a value from channel
	fmt.Println(value)
}
