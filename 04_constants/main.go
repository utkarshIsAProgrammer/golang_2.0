package main

import "fmt"

const (
	port = 5000
	host = "localhost"
)

func main() {
	const name string = "indiedev"
	fmt.Println(name)

	fmt.Println(port, host)
}
