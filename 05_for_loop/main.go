package main

import "fmt"

func main() {
	// while loop like for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// classic for loop
	for i := 1; i <= 7; i++ {
		if i == 2 {
			// break
			continue
		}

		fmt.Println(i)
	}

	// for loop with range
	for i := range 3 {
		fmt.Println(i)
	}

	// infinite loop
	for {
		fmt.Println("infinite loop!")
	}

}
