package main

import "fmt"

type Product struct {
	ID      int
	Name    string
	Price   float64
	InStock bool
}

type Address struct {
	City    string
	Country string
}

type User struct {
	ID      int
	Name    string
	Address Address
}

type Order struct {
	ID         int
	UserID     int
	TotalPrice float64
	Status     string
}

func (p Product) applyDiscount(discount float64) {
	discountedPrice := p.Price - (p.Price * discount / 100)
	fmt.Printf("%s discounted price: %.2f\n", p.Name, discountedPrice)
}

func (p *Product) updatePrice(newPrice float64) {
	p.Price = newPrice
}

func (o Order) isPaid() bool {
	return o.Status == "paid"
}

func (o *Order) markAsPaid() {
	o.Status = "paid"
}

func (o *Order) applyDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		fmt.Println("Invalid discount")
		return
	}

	o.TotalPrice -= o.TotalPrice * discount / 100
}

func (u User) updateNameValue(name string) {
	u.Name = name
	fmt.Println("Inside value receiver:", u.Name)
}

func (u *User) updateName(name string) {
	u.Name = name
}

func main() {
	productOne := Product{
		ID:      2135635,
		Name:    "Laptop",
		Price:   205699,
		InStock: true,
	}

	productTwo := Product{
		ID:      9275675,
		Name:    "TV",
		Price:   15429.56,
		InStock: true,
	}

	productThree := Product{
		ID:      6447449,
		Name:    "SSD",
		Price:   70299,
		InStock: false,
	}

	fmt.Println("\n--- Products ---")
	fmt.Println(productOne)
	fmt.Println(productTwo)
	fmt.Println(productThree)

	fmt.Println("\n--- Product Discount ---")
	productOne.applyDiscount(10)
	fmt.Println("Original price:", productOne.Price)

	fmt.Println("\n--- Update Product Price ---")
	productOne.updatePrice(180000)
	fmt.Println("Updated price:", productOne.Price)

	user := User{
		ID:   53453,
		Name: "Ben 10",
		Address: Address{
			City:    "Nowhere",
			Country: "Austria",
		},
	}

	fmt.Println("\n--- User ---")
	fmt.Println(user)
	fmt.Println(user.Address.City, user.Address.Country)

	order := Order{
		ID:         101,
		UserID:     user.ID,
		TotalPrice: 5000,
		Status:     "pending",
	}

	fmt.Println("\n--- Order ---")

	fmt.Println("Is paid?", order.isPaid())

	// Mark as paid
	order.markAsPaid()
	fmt.Println("Status:", order.Status)
	fmt.Println("Is paid?", order.isPaid())

	order.applyDiscount(20)
	fmt.Printf("Discounted total: %.2f\n", order.TotalPrice)

	fmt.Println("\n--- User Receiver Comparison ---")

	user.updateNameValue("Alex")
	fmt.Println("After value receiver:", user.Name)

	user.updateName("Alex")
	fmt.Println("After pointer receiver:", user.Name)
}
