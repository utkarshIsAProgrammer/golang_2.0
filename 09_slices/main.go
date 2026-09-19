package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil, len(nums))

	// initialized slice with make (default values are 0)
	var nums2 = make([]int, 2, 5)
	fmt.Println(nums2 == nil, len(nums2), cap(nums2))

	// appending elements
	nums2 = append(nums2, 11, 12, 13)
	fmt.Println(nums2, cap(nums2))

	// slice with initial values
	nums3 := []int{16, 7, 9, 19}
	nums3 = append(nums3, 24)
	fmt.Println(nums3, cap(nums3))

	// copy function
	cp1 := []string{"one", "two", "three"}
	cp2 := make([]string, len(cp1))
	copy(cp2, cp1)
	fmt.Println(cp2, cap(cp2))

	// slice operator
	fmt.Println(nums3[1:3])
	fmt.Println(slices.Equal(cp1, cp2))

	// 2-D slice
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix)
	fmt.Println(matrix[0][1])
}

// slice dynamically increase it's size as needed and in start we can specify the length and capacity to add some initial values to make it non-nil
