package main

import "fmt"

func main() {
	fmt.Println("I am main function")

	add(2, 8) // call standerd function
	
	// anonymous function
	add := func(a int, b int) { 
		fmt.Println(a + b)
	} //() invoked immediately

	add(2, 3) // call anonymous (global scope এর add func() কে shadow করে দিয়েছে , এর পর যত add() invoked হবে শুধু anonymous কাজ করবে। )

}

func init() {
	fmt.Println("I am init function and i will invoked first auto")
}
func add(a int, b int) {
	fmt.Println(a + b)
}
