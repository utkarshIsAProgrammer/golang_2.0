package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	product1 := Product{
		ID:    67,
		Name:  "Laptop",
		Price: 67467.89,
	}

	t := reflect.TypeOf(product1)
	v := reflect.ValueOf(product1)

	fmt.Println("Struct Type:", t)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fmt.Println("Field:", field.Name)
		fmt.Println("Type:", field.Type)
		fmt.Println("Value:", v.Field(i))
		fmt.Println("---")
	}
}
