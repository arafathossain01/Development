### Function

Function মানে হলো এমন একটি কোড ব্লক, যেটা নির্দিষ্ট একটা কাজ করার জন্য তৈরি করা হয় এবং প্রয়োজন হলে বারবার ব্যবহার করা যায়।

**Function কেন ব্যবহার করা হয়?**

- একই কাজ বারবার লিখতে হয় না
- কোড ছোট ও পরিষ্কার থাকে
- Debug করা সহজ হয়
- Reusable (পুনরায় ব্যবহারযোগ্য)

Function তখনই কাজ করে যখন তাকে invoked করা হয়।

- Parameter = function এর ভিতরের variable (placeholder)
- Argument = function call করার সময় যে actual value পাঠানো হয়

```go
function add(int a, int b) {   //  a, b => parameters
    return a + b;
}

add(5, 3);  //  5, 3 => arguments

```

```go
func calculation(num1 int, num2 int, num3 int) (sum int, mul int) { // return with identifier
	sum = num1 + num2 + num3
	mul = num1 * num2 * num3
	return // named return
}
```

## Function types:

- **Standard function**

  ```go
  package main
  import "fmt"

  func main() {
  fmt.Println("Hello World")
  }
  ```

- **Init functin (_you can't call it, computer call it auto_)** <a href="./init/test.go">init</a>

  init func() এর কোন input নাই আর এটা কোন কিছু return ও করে নাহ। program শুরু হওয়ার আগে init func() call হবে পরে main func() কল হবে। **init func() --> main func()**

  ```go
  func init(){
  fmt.Println("I'll call first")
  }
  ```

- **Anonymous funcion and IIFE - immediately invoked function expression** <a href="./anonymous/test.go">anonymous</a>

  যে function এর নাম নাই তাই anonymous function. Anonymous function কে কোনো না কোনো function এর ভিতরে রাখতে হবে অথবা variable এ assign করতে হবে। Anonymous function runtime এ তৈরি হয়, memory তে temporary থাকে।

  কোন function কে যদি define করার সাথে সাথে invoked করা হয় তাহলে সেইটা immediately invoked function expression.

  ```go
  func() {
  	fmt.Println("I am Anonymous function")
  }() // invoked immediately
  ```

- **Function expression or Assign function in variable**

  একটা function কে যদি variable-এ assign করা হয়, তাহলে সেটাকে function expression বলা হয়. Function expression এ local scope এ expression এর আগে function invoked করে ফেললে আর কাজ করবে নাহ।

  ```go
  package main
  import "fmt"

  func main(){

  	add (2,3) // not work

  	add := func(a int, b int){ // function expression
  		fmt.Println(a+b)
  	}
  }
  ```

- **Higher order function (first class function) VS first order function**
  - Higher order function following any one of three: <a href="./HOfunction/main.go" >HOfunction</a>

    i. parameter হিসাবে function যাবে<br>
    ii. function return করবে<br>
    iii. both

  - First order functions are simple, they work with:

    i. normal data/ value <br>
    ii. function parameter নেয় না <br>
    iii. function return করে নাহ।

- Callback function

  Higher Order Function এর ভিতরে parameter হিসেবে যে function পাঠানো হয়, সেটাকেই সাধারণভাবে Callback Function বলে।

- First class citizen

  যাদের variable এ assigen করা যায়, function এর argument হিসাবে পাঠানো যায়, function থেকে return করা যায় তাদের first class citizen বলে। যেহেতু go তে function কে ও return করতে পারি, variable এ assigne করতে পারি তাই higher order funtion কে first class function বলা হয়।

- Variadic function

- Closure - close over <a href="./closure/main.go">Colsure</a> <a href="../InternalMemory/note.md" style="color:red">[InternalMemory]</a>
- Defer function - last in first out
- Receiver function <a href="./receiver/main.go">Receiver</a> <a href="../Structure/main.go" style="color:red">[know first struct]</a>
