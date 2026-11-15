package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateValue(value *int) {
	*value = 100
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func updateUser(user *User, newAge int) {
	user.Age = newAge
}

func updateA(user User, newAge int) {
	user.Age = newAge
}

func updateB(user *User, newAge int) {
	user.Age = newAge
}

func main() {
	number := 9
	pointNumber := &number

	fmt.Println("Original value:", number)
	fmt.Println("Pointer address:", pointNumber)
	fmt.Println("Dereferenced value:", *pointNumber)

	updateValue(&number)
	fmt.Println("Updated value:", number)

	x := 10
	y := 12

	swap(&x, &y)
	fmt.Printf("After swap: x = %d | y = %d\n", x, y)

	user := User{
		Name: "John",
		Age:  20,
	}

	updateUser(&user, 25)
	fmt.Println("Updated user:", user)

	var ptr *int

	fmt.Println("Is pointer nil?", ptr == nil)

	if ptr == nil {
		value := 50
		ptr = &value
	}

	fmt.Println("Value through pointer:", *ptr)

	updateA(user, 30)
	fmt.Println("After updateA:", user.Age)

	updateB(&user, 35)
	fmt.Println("After updateB:", user.Age)
}
