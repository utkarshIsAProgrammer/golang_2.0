package main

import "fmt"

/* type payment struct {
	gateway razorpay
	// gateway stripe
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

func main() {
	razorpayGateway := razorpay{}
	// stripeGateway := stripe{}
	p := payment{gateway: razorpayGateway}
	p.makePayment(100.0)
}
*/

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func makeSpeak(s Speaker) {
	s.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeSpeak(dog)
	makeSpeak(cat)
}
