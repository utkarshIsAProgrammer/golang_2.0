package main

import "fmt"

type Logger struct{}

type UserService struct {
	Logger
}

func (l Logger) Log(message string) {
	fmt.Println("Log:", message)
}

func main() {
	service := UserService{}
	service.Log("User created successfully!")
}
