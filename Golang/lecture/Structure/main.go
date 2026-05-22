package main

import "fmt"

type User struct { // struct (class-like data type)
	Name string
	Age  int
}

func main() {

	var user1 User // struct variable / object / instance

	// user1 এর মধ্যে value assign করা হচ্ছে
	user1 = User{
		Name: "Arafat",
		Age:  20,
	}

	fmt.Println("Name: ", user1.Name)
	fmt.Println("Age:  ", user1.Age)

	user2 := User{ // another object / instance
		Name: "Kabila",
		Age:  40,
	}

	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age:  ", user2.Age)
}