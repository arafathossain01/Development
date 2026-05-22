package main
import ("fmt")

func main(){
	var std1 string ;
	std1 = "john"; // assign after declaration

	var num int = 20;
	num2 := 30;

	// num2 := 43;   //❌ error (cannot redeclare in same scope) | Reassign using 	=

	fmt.Println(std1);
	fmt.Println(num);
	fmt.Println(num2);

}

/*
	In Go programming language, variables can be declared in two main ways. Go is strict and clear about how variables are declared and assigned.
	1. using `var`
	2. using `:=`

	num2 := 30;
	- Type is automatically inferred
	- Only works inside functions
	- Must assign value at the same time
*/