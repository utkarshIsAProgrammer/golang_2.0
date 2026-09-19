package main

import "fmt"

func changeNum(a *int) {
	*a = 5
	fmt.Println(*a)
}

func main() {
	num := 1

	changeNum(&num)
	fmt.Println("After change:", num)
}
