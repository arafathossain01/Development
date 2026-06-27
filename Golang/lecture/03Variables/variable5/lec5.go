package main
import ("fmt")

const PI = 3.1416; // untyped
const num int = 40; //typed

func main(){

	const (
		num1 int = 10;
		num2 = 1.45;
		str string = "Hello World";
	)
	fmt.Println(num)
	fmt.Println(PI)
	fmt.Println(str)
}

/*
	The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.  

	The value of a constant must be assigned when declare it.
	
	Constants can be declared both inside and outside of a function
*/
