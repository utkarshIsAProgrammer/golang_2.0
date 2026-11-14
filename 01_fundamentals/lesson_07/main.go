package main

import "fmt"

type Student struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Price float64
}

type StudentData struct {
	Subject string
	Marks   int
}

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func countWords(words []string) map[string]int {
	result := make(map[string]int)

	for _, word := range words {
		result[word]++
	}

	return result
}

func findUser(users map[int]User, id int) (User, bool) {
	user, ok := users[id]
	return user, ok
}

func main() {
	students := map[int]Student{
		1: {ID: 1, Name: "Alex"},
		2: {ID: 2, Name: "Peter"},
	}

	students[3] = Student{ID: 3, Name: "Bob"}
	students[4] = Student{ID: 4, Name: "Charlie"}
	students[5] = Student{ID: 5, Name: "Kevin"}

	fmt.Println("Student:", students[3])

	students[3] = Student{ID: 3, Name: "Diva"}

	delete(students, 2)

	fmt.Println("Updated student:", students[3])
	fmt.Println("Total students:", len(students))

	products := map[string]Product{
		"Laptop":  {ID: 23, Price: 95799},
		"Mobile":  {ID: 14, Price: 42499.63},
		"Monitor": {ID: 82, Price: 105000},
		"Ram":     {ID: 27, Price: 72564.12},
	}

	products["SSD"] = Product{ID: 9, Price: 36500.48}

	fmt.Println("\nProduct:", products["Ram"])

	result, ok := products["Rams"]

	if ok {
		fmt.Println("Product ID:", result.ID)
		fmt.Println("Product Price:", result.Price)
	} else {
		fmt.Println("Product not found")
	}

	sol := countWords([]string{
		"go", "rust", "xaml", "go", "go", "rust",
	})

	fmt.Println("\nWord count:", sol)

	marks := map[string]StudentData{
		"Jane": {
			Subject: "Math",
			Marks:   54,
		},
		"John": {
			Subject: "English",
			Marks:   42,
		},
	}

	marks["David"] = StudentData{
		Subject: "Science",
		Marks:   78,
	}

	fmt.Println("\nDavid's subject:", marks["David"].Subject)
	fmt.Println("David's marks:", marks["David"].Marks)

	john := marks["John"]
	john.Marks = 90
	marks["John"] = john

	fmt.Println("Updated John's marks:", marks["John"].Marks)

	alice, exists := marks["Alice"]

	if exists {
		fmt.Println("Alice:", alice)
	} else {
		fmt.Println("Alice not found")
	}

	fmt.Println("Total students in marks map:", len(marks))

	users := map[int]User{
		1: {ID: 1, Name: "Alex", Email: "alex@example.com", Age: 21},
		2: {ID: 2, Name: "Peter", Email: "peter@example.com", Age: 22},
		3: {ID: 3, Name: "Kevin", Email: "kevin@example.com", Age: 23},
	}

	user, found := findUser(users, 2)

	if found {
		fmt.Println("\nUser found:", user)
	} else {
		fmt.Println("\nUser not found")
	}
}
