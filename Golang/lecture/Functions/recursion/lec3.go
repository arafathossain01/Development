package main

import "fmt"

func count(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return count(x + 1)
}
func main() {
	count(1)
}

/*
Recursion is a process where a function calls itself repeatedly until a base condition is satisfied
*/
