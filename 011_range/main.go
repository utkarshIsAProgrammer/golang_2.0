package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	var sum int

	for idx, num := range nums {
		fmt.Println(idx+1, num)
		sum += num
	}
	fmt.Println(sum)

	m := map[string]string{
		"fname": "John",
		"lname": "Doe",
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
