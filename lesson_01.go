package main

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type (
	CardPayment struct{}
	UPIPayment  struct{}
)

func (c CardPayment) Pay(amount float64) {
	fmt.Println("Paid", amount, "using Card")
}

func (u UPIPayment) Pay(amount float64) {
	fmt.Println("Paid", amount, "using UPI")
}

func processPayment(p Payment, amount float64) {
	p.Pay(amount)
}

func main() {
	card := CardPayment{}
	upi := UPIPayment{}

	processPayment(card, 500)
	processPayment(upi, 1000)
}
