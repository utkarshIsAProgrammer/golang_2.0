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

func countWords(words []string) map[string]int {
	result := make(map[string]int)

	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	students := map[int]Student{
		1: {ID: 01, Name: "Alex"},
		2: {ID: 02, Name: "Peter"},
	}

	students[3] = Student{ID: 03, Name: "Bob"}
	students[4] = Student{ID: 04, Name: "Charlie"}
	students[5] = Student{ID: 05, Name: "Kevin"}

	fmt.Println(students[03])

	students[3] = Student{ID: 03, Name: "Diva"}
	delete(students, 02)

	fmt.Println("Updated student:", students[03])
	fmt.Println("Totla students:", len(students))

	products := map[string]Product{
		"Laptop":  {ID: 23, Price: 95799},
		"Mobile":  {ID: 14, Price: 42499.63},
		"Monitor": {ID: 82, Price: 105000.00},
		"Ram":     {ID: 27, Price: 72564.12},
	}

	products["SSD"] = Product{ID: 9, Price: 36500.48}

	fmt.Println(products["Ram"])

	result, ok := products["Rams"]
	if ok {
		fmt.Println("Product ID:", result.ID)
		fmt.Println("Product Price:", result.Price)
	} else {
		fmt.Println("Product not found")
	}

	sol := countWords([]string{"go", "rust", "xaml", "go", "go", "rust"})
	fmt.Println(sol)

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

	fmt.Println(marks["Jane"].Subject, marks["Jane"].Marks)
}
