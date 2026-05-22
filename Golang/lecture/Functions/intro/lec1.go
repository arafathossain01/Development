package main

import "fmt"

func add(num1 int, num2 int) {
	var result int = num1 + num2
	fmt.Println("The sum:", result)
}
func main() {
	var (
		number1 int
		number2 int
	)
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&number1, &number2)

	add(number1, number2)

}
