package main

import "fmt"

func call() func(a int, b int) { // return function
	return addad
}

func addad(a int, b int) {
	z := a + b
	fmt.Println(z)
}
func main() {
	sum := call() // return করার পর sum এর মধ্যে add আসবে
	sum(4, 3)     // add কে call করা হচ্ছে
}
