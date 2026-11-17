package main

import "fmt"

func main() {
	fmt.Println("First ")
	fmt.Println("Second")
	defer fmt.Println("Third")
	defer fmt.Println("Fourth")

	panic("error occurred!")
}
