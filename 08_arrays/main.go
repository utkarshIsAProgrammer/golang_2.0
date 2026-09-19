package main

import "fmt"

func main() {
	// making array (by default values are 0)
	var nums [5]int
	nums[0] = 1
	nums[1] = 2

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(nums[2], nums[0])

	// making array (by default values are false)
	var vals [4]bool
	vals[0] = true
	vals[1] = false

	fmt.Println(vals)

	// making array (by default values are "")
	var names [3]string
	names[2] = "indiedev"

	fmt.Println(names)

	// single line array initialization
	ages := [3]int{28, 21, 26}
	fmt.Println(ages)

	// 2-D array initialization
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)
}

// fixed size, that is predictable
// memory optimization
// fast time access
