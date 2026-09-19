package main

import "fmt"

func main() {
	// create a map
	m := make(map[string]any)

	// setting elements
	m["name"] = "indiedev"
	m["age"] = 20

	fmt.Println(m)
	fmt.Println(m["name"], m["age"])
	fmt.Println(len(m))

	// delete element
	delete(m, "age")
	fmt.Println(m)

	// clear map
	clear(m)
	fmt.Println(m)

	// map without make
	n := map[string]any{"name": "indiedev", "age": 20, "admin": true, "role": "backend developer"}
	fmt.Println(n)
}
