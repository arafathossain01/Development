package main

import "fmt"

var arr2 = [3]string{"I", "Love", "You"}

func main() {
	// var arr [2] int
	// arr[0] = 1
	// arr[1]= 2

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr2)

	arr3 := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < 5; i++ {
		fmt.Println(arr3[i])
	}
}

/*
	Array index start from 0
*/
