package main

import "fmt"

func main() {
	var day int

	fmt.Print("Enter the day: ")
	fmt.Scan(&day)

	switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Wrong input!")


	// multi-case
	
	case 1, 3, 5:
		fmt.Println("odd weekday")
	case 2,4:
		fmt.Println("Even weekday")
	case 6,7:
		fmt.Println("weekend")
	default:
		fmt.Println("Invalid input")

	}
}

// switch case is faster than if else.

/*
	All the case values should have the same type as the switch expression.

*/