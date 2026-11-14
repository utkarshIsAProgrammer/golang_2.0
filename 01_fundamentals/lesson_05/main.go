package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func analyzeNumber(number int) (bool, bool) {
	return number%2 == 0, number > 0
}

func calculateAverage(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("no numbers provided")
	}

	var total float64

	for _, value := range numbers {
		total += value
	}

	return total / float64(len(numbers)), nil
}

func validateCredentials(username, password string) bool {
	return username == "admin" && password == "admin123"
}

func checkAccountStatus(isBlocked bool) bool {
	return !isBlocked
}

func login(username, password string, isBlocked bool) (string, error) {
	if !validateCredentials(username, password) {
		return "", errors.New("invalid credentials")
	}

	if !checkAccountStatus(isBlocked) {
		return "", errors.New("account is blocked")
	}

	return "Login Successful!", nil
}

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	a := add(9, 7)
	b := subtract(9, 7)
	c := multiply(5, 6)

	fmt.Printf("Sum: %d\n", a)
	fmt.Printf("Subtract: %d\n", b)
	fmt.Printf("Multiply: %d\n", c)

	d, err := divide(30, 2)

	if err != nil {
		fmt.Println("Divide:", err)
	} else {
		fmt.Printf("Divide: %d\n", d)
	}

	even, positive := analyzeNumber(0)

	fmt.Printf("Even: %t\n", even)
	fmt.Printf("Positive: %t\n", positive)

	numbers := []float64{99, 45, 64, 82, 91, 29, 85, 64, 73, 72}

	average, err := calculateAverage(numbers...)

	if err != nil {
		fmt.Println("Average error:", err)
	} else {
		fmt.Printf("Average: %.2f\n", average)
	}

	fmt.Println("Valid credentials:", validateCredentials("admin", "admin123"))
	fmt.Println("Account not blocked:", checkAccountStatus(false))

	message, err := login("admin", "admin123", false)

	if err != nil {
		fmt.Println("Login error:", err)
	} else {
		fmt.Println(message)
	}

	myCounter := counter()

	fmt.Println("Counter:", myCounter())
	fmt.Println("Counter:", myCounter())
	fmt.Println("Counter:", myCounter())
	fmt.Println("Counter:", myCounter())
}
