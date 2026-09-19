package main

import "fmt"

func main() {
	// if, else if and else
	age := 15

	if age >= 18 {
		fmt.Println("You're an adult!")
	} else if age >= 12 {
		fmt.Println("You're a teenager!")
	} else {
		fmt.Println("You're a minor!")
	}

	// if with comparison operator
	role := "admin"
	hasPermissions := true

	if role == "admin" || hasPermissions {
		fmt.Println("You have admin permissions!")
	} else {
		fmt.Println("You don't have admin permissions!")
	}

	// var declaration and if condition
	if age := 12; age >= 18 {
		fmt.Println("You're an adult!")
	} else {
		fmt.Println("You're a minor!")
	}
}
