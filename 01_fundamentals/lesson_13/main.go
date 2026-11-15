package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

var ErrInsufficientBalance = errors.New("insufficient balance")

func withdraw(balance, amount float64) (float64, error) {
	if amount < 0 {
		return balance, errors.New("withdrawal amount cannot be negative")
	}

	if amount > balance {
		return balance, ErrInsufficientBalance
	}

	return balance - amount, nil
}

var ErrUserNotFound = errors.New("user not found")

func fetchUser(id int) error {
	return ErrUserNotFound
}

func getUser(id int) error {
	err := fetchUser(id)
	if err != nil {
		return fmt.Errorf("failed to get user with ID %d: %w", id, err)
	}

	return nil
}

func main() {
	fmt.Println("--- Division ---")

	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("\n--- Bank Withdrawal ---")

	balance := 1000.0

	newBalance, err := withdraw(balance, 300)

	if err != nil {
		if errors.Is(err, ErrInsufficientBalance) {
			fmt.Println("Error: Insufficient balance")
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		balance = newBalance
		fmt.Println("Remaining balance:", balance)
	}

	newBalance, err = withdraw(balance, 2000)

	if err != nil {
		if errors.Is(err, ErrInsufficientBalance) {
			fmt.Println("Error: Insufficient balance")
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		balance = newBalance
		fmt.Println("Remaining balance:", balance)
	}

	fmt.Println("\n--- Fetch User ---")

	err = getUser(101)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("Original error identified:", ErrUserNotFound)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
