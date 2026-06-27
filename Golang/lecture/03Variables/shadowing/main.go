package main

import "fmt"

var a = 10 // global variable

func main() {
	age := 30

	if age > 18 {
		a := 47 // declare new variable and shadowing here  | output: 47
		fmt.Println(a)
	}

	fmt.Println(a) // output: 10
}
