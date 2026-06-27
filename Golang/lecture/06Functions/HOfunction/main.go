package main

import "fmt"

func processOperation(x int, y int, OP func(a int, b int)) { // OP func (a int, b int) => callback function
	OP(x, y)
}

func add(a int, b int) {
	z := a + b
	fmt.Println(z)
}

func main() {
	processOperation(4, 5, add)
}

/*
প্রথমে main() function run হবে।

তারপর:

processOperation(4, 5, add)

call হবে।

এখানে:
4 যাবে x এর ভিতরে
5 যাবে y এর ভিতরে
add function যাবে OP এর ভিতরে

অর্থাৎ:

x = 4
y = 5
OP = add

এরপর processOperation() এর ভিতরে:

OP(x, y)

execute হবে।

যেহেতু OP এর ভিতরে add function আছে,
তাই এটা আসলে:

add(4, 5)

হিসাবে কাজ করবে।

তারপর add() function এ:

a = 4
b = 5

হবে।

এরপর:

z := a + b

মানে:

z = 4 + 5
z = 9

তারপর:

fmt.Println(z)

9 print করবে।
*/