package main

import "fmt"

// [T any] means T is a type parameter that can be any type
// (value T) means value is of type T
func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	printValue(42)
	printValue("hello")
	printValue(true)
	printValue(3.14)
}
