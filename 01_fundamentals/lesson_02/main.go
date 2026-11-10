package main

import "fmt"

func main() {
	var name string
	var age int
	var email string
	var isVerified bool
	var accountBalance float32

	const (
		ApplicationName    = "GoBank"
		MaxLoginAttempts   = 5
		ApplicationVersion = "1.0.0"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(email)
	fmt.Println(isVerified)
	fmt.Println(accountBalance)

	fmt.Println(ApplicationName)
	fmt.Println(MaxLoginAttempts)
	fmt.Println(ApplicationVersion)
}
