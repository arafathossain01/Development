package main
import ("fmt")

func main(){
	var a, b, c, d int = 1, 2, 3,  4; // this is multi variable declaration on one line


	/*
		If use varibale type, then it is possible to declare one type of variable per line.
	*/

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var  num, str = 5, "hello";

	/*
		If not specified the variable type, then declare different types  of variable on the same line.
	*/

	fmt.Println(num);
	fmt.Println(str)
}
