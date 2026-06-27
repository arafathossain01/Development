package main
import ("fmt")

func main(){
	var(
		num1 int = 20;
		num2 int = 30;
		fName string = "Sabikun";
		lName string = "Nahar";
	)

	fmt.Println(num1) // print with new line 
	fmt.Println(num2)
	fmt.Print(fName, " ") // print without new line
	fmt.Print(lName, "\n")

	var i string = "Hello";
	var j int = 20;
	fmt.Printf("i has value: %v and type: %T\n", i, i) // %T means the type & %v means the value.
	fmt.Printf("j has value: %v and type: %T\n", j, i)
}

/*
	`\n`--> is new line character
	" " --> add space 
	The fmt.Println() function adds a space between multiple values and also prints a new line at the end.
*/