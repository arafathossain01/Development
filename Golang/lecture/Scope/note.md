### What is Scope?

Go-তে Scope মানে হলো কোন variable বা function কোথায় use করা যাবে

Scope are three types:

1. Global Scope

   Go এর package এর সব জায়গায় থেকে access করা যাবে একে। function এর বাইরে declare করা হয়। Top-level function গুলা global scope এ add হবে। 

2. Local Scope (function scope)

   শুধু function এর ভিতরে কাজ করবে variable. function এর ভিতরে declare করা হয়। একটা function এর ভিতর অন্য একটা function থাকলে ওইটা local scope এ থাকবে। 

3. Block Scope

   `{}` blcok scope. block scope এর ভিতরে variable declare করলে variable শুধু ওই block এর ভিতরেই কাজ করবে, বাইরে কাজ করবে নাহ।

```go
package main
import "fmt"

var a = 100 // global

func main() {
    b := 20 // local

    if true {
        c := 10 // block
        fmt.Println(a, b, c)
    }

    // fmt.Println(c) ❌ error
}
```

**যখন নিজের scope এ খুঁজে না পায় value তখন global scope এ যায় program**

<h1> Package Scope </h1>

Package scope হচ্ছে একটা package থেকে অন্য একটা package এর জিনিস পত্র ব্যবহার করা। package মানে একটা folder. package এর নাম হয় তার folder এর নাম অনুসারে।

একই folder এ একাধিক file থাকলে package এর নাম same দেওয়া লাগবে। একাধিক file একসাথে run করতে হলে: **go run file1.go file2.go file3.go.....**

<h3> Package এর নাড়িভুঁড়ি</h3>

Go-তে অন্য package থেকে code ব্যবহার করতে হলে প্রথমে একটি module তৈরি করতে হয়। এজন্য `go mod init module_name` কমান্ড ব্যবহার করা হয়। এর মাধ্যমে project-এর জন্য একটি module setup করা হয়, যা dependency management সহজ করে।

অন্য package ব্যবহার করতে হলে সেটি `import` করতে হয়। এরপর সেই package-এর function বা variable ব্যবহার করা যায়।

তবে একটি গুরুত্বপূর্ণ নিয়ম হলো—অন্য package থেকে কোনো function, variable বা struct ব্যবহার করতে চাইলে সেগুলোর নাম অবশ্যই **Capital letter (বড় হাতের অক্ষর)** দিয়ে শুরু হতে হবে। কারণ Go-তে capital letter দিয়ে শুরু হওয়া identifier গুলো exported (public) হয় এবং অন্য package থেকে access করা যায়। আর small letter দিয়ে শুরু হলে তা private থাকে এবং শুধুমাত্র ওই package-এর মধ্যেই ব্যবহার করা যায়।

সাধারণভাবে, main package হলো program-এর entry point, যেখানে অন্য package থেকে প্রয়োজনীয় function import করে ব্যবহার করা হয়।
