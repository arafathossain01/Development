package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter marks: ")
	fmt.Scan(&num)

	if num >= 80 && num <= 100 {
		fmt.Println("You got A+")
	} else if num >= 70 && num <= 79 {
		fmt.Println("You got A")
	} else if num >= 60 && num <= 69 {
		fmt.Println("You got B+")
	} else if num >= 50 && num <= 59 {
		fmt.Println("You got B")
	} else if num >= 40 && num <= 49 {
		fmt.Println("You got C")
	} else if num >= 33 && num <= 39 {
		fmt.Println("You got D")
	} else {
		fmt.Println("You are Fail")
	}
}

/*
	Having the else brackets in a different line will raise an error

	if (temperature > 15) {
    	fmt.Println("It is warm out there.")
  	} // this raises an error
  	else {
    	fmt.Println("It is cold out there.")
  	}
*/
