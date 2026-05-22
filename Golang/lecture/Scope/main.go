package main

import (
	"example.com/mathlib" // custom package import
	"fmt"
)

var (
	a = 10 // global scope
	b = 20
)

func main() {
	var c = 30 // local scope
	var d = 50

	add(c, d)
	add(a, b)
	fmt.Print("Multiplication from custom package: ")
	mathlib.Mul(a, b)
	if a == 10 {
		fmt.Println("I am", a, "years old")
		d := 40 // block scope
		fmt.Print(d)
	}
}
