package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age: ", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}
func init() {
	fmt.Println("===Start===")
}


/*
এখানে Go প্রথমে program এর code গুলো Code Segment এ রাখবে।

যেমন:
- constant a
- outer()
- anonymous function (show)
- call()
- main()
- init()

তারপর runtime এ Data Segment এ global variable p থাকবে।

কারণ:
var p = 100

এটা global variable।

------------------------------------------------

Program run হওয়ার সময় প্রথমে init() function execute হবে।

তখন output হবে:

===Start===

এরপর main() function run হবে।

main() এর ভিতরে call() execute হবে।

------------------------------------------------

call() এর ভিতরে:

incr1 := outer()

লাইন execute হলে outer() function call হবে।

তখন Stack এ outer() এর জন্য একটি Stack Frame তৈরি হবে।

Stack Frame এ থাকবে:

money = 100
age = 30

তারপর:

fmt.Println("Age: ", age)

execute হবে।

Output:

Age: 30

------------------------------------------------

এরপর show নামে একটি anonymous function তৈরি হবে।

এই function এর ভিতরে:
- money (local variable)
- a (constant)
- p (global variable)

ব্যবহার হচ্ছে।

show function money variable কে মনে রাখবে।
এটাকে Closure বলে।

------------------------------------------------

এরপর:

return show

হবে।

মানে show function return হবে।

এখন:

incr1 = show

------------------------------------------------

Normally outer() function শেষ হলে
তার Stack Frame remove হয়ে যেত।

তাহলে:
money variable ও destroy হয়ে যেত।

কিন্তু এখানে show function পরে money ব্যবহার করবে।

তাই Go compiler বুঝতে পারে:

"money variable function শেষ হওয়ার পরও দরকার হবে"

তখন money Stack থেকে Heap এ যেতে পারে।

এটাকে Escape Analysis বলে।

Heap এ যাওয়ার কারণে:
money variable outer() শেষ হওয়ার পরও alive থাকে।

Heap memory Go এর Garbage Collector (GC)
manage করে।

------------------------------------------------

তারপর:

incr1()

call হবে।

এখন show function execute হবে।

তখন:

money = money + a + p

মানে:

money = 100 + 10 + 100
money = 210

Output:

210

------------------------------------------------

আবার:

incr1()

call হলে আগের money value retain থাকবে।

কারণ closure money কে মনে রেখেছে
এবং money Heap এ alive আছে।

এখন:

money = 210 + 10 + 100
money = 320

Output:

320

------------------------------------------------

এরপর:

incr2 := outer()

আবার নতুন করে outer() call হবে।

তখন নতুন Stack Frame তৈরি হবে।

নতুন:
money = 100
age = 30

তৈরি হবে।

আবার Output:

Age: 30

------------------------------------------------

তারপর:

incr2()

হলে:

money = 210

print হবে।

আবার:

incr2()

হলে:

money = 320

print হবে।

------------------------------------------------

এখানে incr1 এবং incr2 আলাদা closure।

তাই:
- incr1 এর money আলাদা
- incr2 এর money আলাদা

------------------------------------------------

Final Output:

===Start===
Age: 30
210
320
Age: 30
210
320
*/