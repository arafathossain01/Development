package main

import "fmt"

func message() {
	fmt.Println("----------Welcome to the information saver----------")
}

func userName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var (
		num1 int
		num2 int
	)
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	var sum int = num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("Your name is: ", name)
	fmt.Println("Your sum is: ", sum)
}

func goodBye(name string) {
	fmt.Println("GoodBye ", name)
}

func main() {
	message()
	var name string = userName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	display(name, sum)
	goodBye(name)
}
