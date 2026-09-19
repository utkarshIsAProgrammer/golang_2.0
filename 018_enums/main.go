package main

import "fmt"

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

type Value int

const (
	Zero Value = iota
	One
	Two
	Three
)

func main() {
	var status Status = Approved
	fmt.Println(status)

	var value Value = Two
	fmt.Println(value)
}
