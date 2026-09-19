package main

import "fmt"

func sum(num ...int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
}
