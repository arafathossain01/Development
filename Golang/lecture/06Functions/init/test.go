package main

import "fmt"

var a = 10

func main() {
	fmt.Println(a) // 20
}

func init() {
	fmt.Println(a) // 10
	a = 20
}

/*
এই প্রোগ্রামে প্রথমে global scope-এ a = 10 initialize করা হয়।

এরপর Go language-এর execution rule অনুযায়ী main() function run হওয়ার আগে init() function execute হয়।

init() function-এর ভিতরে যখন fmt.Println(a) লেখা হয়, তখনো a এর মান change হয় নাই। তাই global scope থেকে a = 10 নিয়ে print করে।

এরপর a = 20 assign করা হয়, যা global variable a-এর মান update করে।

তারপর main() function execute হয়। তখন a এর updated value 20 হয়ে গেছে, তাই fmt.Println(a) থেকে 20 print হয়।
*/
