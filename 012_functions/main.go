package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "python", "javascript"
}

func processIt(fn func(a int) int) {
	result := fn(7)
	fmt.Println(result)
}

func main() {
	result := add(2, 9)
	fmt.Println(result)

	l1, l2, l3 := getLanguages()
	fmt.Println(l1, l2, l3)

	fn := func(a int) int {
		return a
	}
	processIt(fn)
}
