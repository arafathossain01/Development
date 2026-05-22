package main

import (
	"fmt"
)

func main() {
	var (
		num1 int
		num2 int
		num3 int
	)

	fmt.Print("Enter three numbers: ")
	fmt.Scan(&num1, &num2, &num3)

	if num1 >= num2 && num1 >= num3 {
		fmt.Printf(" %v is bigger than %v and %v", num1, num2, num3)
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Printf(" %v is bigger than %v and %v", num2, num1, num3)
	} else {
		fmt.Printf(" %v is bigger than %v and %v", num3, num1, num2)
	}

}
