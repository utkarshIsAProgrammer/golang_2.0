package main

import "fmt"

func generateNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go generateNumbers(ch)

	for i := 0; i < 5; i++ {
		val := <-ch
		fmt.Println(val)
	}

}
