package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Sunday"

	switch day {
	case "Monday":
		fmt.Printf("%s its Weekday!\n", day)
	case "Tuesday":
		fmt.Printf("%s its Weekday!\n", day)
	case "Wednesday":
		fmt.Printf("%s its Weekday!\n", day)
	case "Thursday":
		fmt.Printf("%s its Weekday!\n", day)
	case "Friday":
		fmt.Printf("%s its Weekday!\n", day)
	case "Saturday":
		fmt.Printf("%s its Weekend!\n", day)
	case "Sunday":
		fmt.Printf("%s its Weekend!\n", day)
	default:
		fmt.Println("Invalid day!")
		fmt.Println(time.Now().Weekday())
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("unknown")
		}
	}

	whoAmI("nine")
}
