package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}

func main() {
	customer1 := customer{
		name:  "devi prasad",
		phone: "9348293467",
	}

	order1 := order{
		id:        "01",
		amount:    100.00,
		status:    "ordered",
		createdAt: time.Now(),
		customer:  customer1,
	}

	order2 := order{
		id:        "02",
		amount:    75.25,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println(order1)
	fmt.Println(order1.status)
	fmt.Println(order1.customer)

	fmt.Println(order2)
	fmt.Println(order2.status)

	order1.changeStatus("packed")
	fmt.Println(order1.status)

	order3 := newOrder("03", 150.50, "shipped")
	fmt.Println(order3)
	fmt.Println(order3.status)

	// struct literals
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
