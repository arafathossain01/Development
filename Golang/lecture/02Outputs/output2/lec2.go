package main
import "fmt"

func main(){
	fmt.Printf("%T\n", 30)
	fmt.Printf("%T\n", "sijuka")
}

/*
	Formatting verbs for printf()
	%v --> default format (সব ধরনের data এর জন্য)
	%d --> integer number print করার জন্য
	%s --> string print করার জন্য
	%f --> decimal print (%.2f) == 3.24
	%T --> variable এর type দেখায়
	%t --> true/false
	%p --> memory address
*/