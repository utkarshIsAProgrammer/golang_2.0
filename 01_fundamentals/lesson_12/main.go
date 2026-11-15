package main

import (
	"fmt"

	"github.com/UtkarshIsAProgrammer/go-modules/calculator"
	"github.com/UtkarshIsAProgrammer/go-modules/user"
)

func main() {
	u := user.User{
		ID:   1,
		Name: "Alex",
		Age:  22,
	}

	fmt.Println("User:", u)

	fmt.Println("Addition:", calculator.Add(10, 5))
	fmt.Println("Subtraction:", calculator.Subtract(10, 5))
	fmt.Println("Multiplication:", calculator.Multiply(10, 5))
}

