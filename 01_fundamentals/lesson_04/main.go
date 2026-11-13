package main

import (
	"fmt"
)

func main() {
	number := 2

	switch {
	case number < 0:
		fmt.Println("It's a negative number.")
	case number > 0:
		fmt.Println("It's a positive number.")
	default:
		fmt.Println("It's a zero.")
	}

	if number%2 == 0 {
		fmt.Println("It's even.")
	} else {
		fmt.Println("It's odd.")
	}

	var marks int32 = 82
	if marks < 0 || marks > 100 {
		fmt.Println("Invalid marks.")
	} else if marks >= 90 {
		fmt.Println("A")
	} else if marks >= 80 {
		fmt.Println("B")
	} else if marks >= 70 {
		fmt.Println("C")
	} else if marks >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("Fail")
	}

	number = 4
	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

	username := "admin"
	password := "secret123"
	isBlocked := true

	if username == "admin" && password == "secret123" {
		if isBlocked {
			fmt.Println("Account blocked!")
		} else {
			fmt.Println("Login successful!")
		}
	} else {
		fmt.Println("Invalid credentials!")
	}

	var total int
	var processedOrders []int
	var highOrderValues []int

	orders := []int{500, 1200, 800, 2000, 300}
	for _, value := range orders {
		if value < 500 {
			continue
		}

		processedOrders = append(processedOrders, value)
		total += value

		if value >= 1000 {
			highOrderValues = append(highOrderValues, value)
		}
	}

	fmt.Println("Processed orders:", processedOrders)
	fmt.Println("High value orders:", highOrderValues)
	fmt.Println("Orders processed:", len(processedOrders))
	fmt.Println("Total amount:", total)
}
