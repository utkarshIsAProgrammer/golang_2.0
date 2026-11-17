package main

import "fmt"

func findMax[T int | float64](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	result := findMax(20, 45)
	fmt.Println(result)
}
