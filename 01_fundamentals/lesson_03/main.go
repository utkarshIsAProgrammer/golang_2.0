package main

import (
	"fmt"
)

func main() {
	a := 25
	b := 4

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	age := 20
	isVerified := true
	isBlocked := false

	if age >= 18 {
		fmt.Println("Users is atleast 18.")
	}

	if isVerified {
		fmt.Println(isVerified)
	}

	if isVerified && !isBlocked {
		fmt.Println("You can access the application.")
	}

	if age < 18 || isBlocked {
		fmt.Println("Either under 18 or blocked.")
	}

	productPrice := 800.0
	quantity := 2.0
	total := productPrice * quantity

	if isVerified && !isBlocked {
		if total >= 1500.0 {
			discount := 10.0
			discountAmount := (total * discount) / 100.0
			finalAmount := total - discountAmount
			if finalAmount >= 1000.0 {
				deliveryFee := 0.0
				finalAmount += deliveryFee
			}
			fmt.Printf("Total: %.2f Discount: %.2f Final Amount: %.2f Verified: %t Blocked: %t\n", total, discount, finalAmount, isVerified, isBlocked)
		}
	}
}
