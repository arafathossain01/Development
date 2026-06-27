package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age:  ", usr.Age)
}

// এটা একটি method (receiver function)  User struct এর সাথে bind করা হয়েছে
func (usr User) printDetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age:  ", usr.Age)
}

// আরেকটি method, এখানে parameter নেওয়া হয়েছে (a int)
func (usr User) call(a int) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age:  ", usr.Age)
	fmt.Println("Number: ", a)
}

func main() {
	user1 := User{
		Name: "Arafat",
		Age:  23,
	}
	// printUserDetails(user1)
	user1.printDetails()

	user2 := User{
		Name: "Tarek",
		Age:  32,
	}
	printUserDetails(user2)
	user2.call(20)
}
