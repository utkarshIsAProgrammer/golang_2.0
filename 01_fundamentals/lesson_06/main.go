package main

import "fmt"

func filterEvenNumbers(numbers []int) []int {
	result := []int{}

	for _, i := range numbers {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	var total int
	largest := 0
	for _, value := range numbers {
		fmt.Println(value)
		total += value

		if largest < value {
			largest = value
		}
	}
	fmt.Println(total)
	fmt.Println(largest)

	numSlice := []int{11, 12, 13, 14, 15}
	numSlice = append(numSlice, 16, 17, 18)
	fmt.Println(len(numSlice), cap(numSlice))

	newNumSlice := numSlice[2:5]
	newNumSlice[1] = 99

	fmt.Println(newNumSlice)
	fmt.Println(numSlice)

	makeSlice := make([]int, 3, 5)
	fmt.Printf("Length before: %d | Capacity before: %d\n", len(makeSlice), cap(makeSlice))

	makeSlice = append(makeSlice, 97, 88)
	makeSlice = append(makeSlice, 12)
	fmt.Printf("Length after: %d | Capacity after: %d\n", len(makeSlice), cap(makeSlice))

	filtered := filterEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("Filtered Evenly:", filtered)
}
