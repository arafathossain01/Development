package main
import ("fmt")

func main()  {
	var  a string;
	var b int 
	var c bool

	fmt.Println(a) // ""
	fmt.Println(b) // 0
	fmt.Println(c) // false
}

/*
	In Go programming language, when a variable is declared but not assigned any value, Go automatically assigns a default value.
	
	- Go does not allow uninitialized variables
	- Every variable gets a zero value automatically
	- No garbage/undefined values like C/C++
	- Helps avoid runtime errors

	Common zero value
	int → 0
	float → 0.0
	string → ""
	bool → false
	pointer → nil
*/