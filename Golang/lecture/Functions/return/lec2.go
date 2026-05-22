package main

import "fmt"

// return one thing
func sum(num1 int, num2 int) int {
	return num1 + num2

}

// return multiple thing
func calculation(num1 int, num2 int, num3 int) (int, int) {
	sum := num1 + num2 + num3
	mul := num1 * num2 * num3
	return sum, mul

}
func main() {
	var n int = 10
	var m int = 20
	var o int = 5

	sum, mul := calculation(n, m, o)
	fmt.Println(sum, mul)
}
