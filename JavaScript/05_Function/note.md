### function statement or function declaration

function keyword দিয়ে যখন কোন function create করা হয় তখন তাকে function statement অথবা function declaration বলে।

```js
function print() {
  console.log("Hello World");
}
```

### function expression or named function expression

যখন কোন variable এর ভিতরে function রাখা হয় তখন তাকে function expression বলে। function expresion creation phase এ memory component এ undefined বসবে এবং execition phase এ পুড়া function টা বসে যাবে code component এ।

```js
var log = function print() {
  console.log("This is function expression");
};
```

### funcion defination vs function expression

- **Function Declaration (Definition):** হোইস্টিং মেকানিজম: এটি হোইস্টিং হওয়ার সময় কোনো এরর দেয় না। কারণ Creation Phase-এই গ্লোবাল মেমরি কম্পোনেন্টে পুরো ফাংশনের বডিটি (f()) হুবহু স্টোর হয়ে থাকে। ফলে ডিক্লেয়ার করার আগেই একে সফলভাবে কল করা যায়।

- **Function Expression:** <a href="../Error/note.md#typeerror">TypeError</a>
  - **আগে কল করলে (TypeError):** ডিক্লেয়ার বা অ্যাসাইন করার আগে কল করলে এটি TypeError দেয়। কারণ হোইস্টিং-এর কারণে এটি শুরুতে মেমরি কম্পোনেন্টে undefined হয়ে থাকে। কোড কম্পোনেন্ট যখন ডিক্লেয়ারেশনের আগেই একে ফাংশন হিসেবে কল করতে চায়, তখন undefined টাইপের সাথে ফাংশন টাইপ ম্যাচ না করায় ইঞ্জিন এরর ছুড়ে দেয়।

  - **পরে কল করলে (সফল এক্সিকিউশন):** কিন্তু এক্সপ্রেশনের লাইনের পরে কল করলে আর কোনো এরর হয় না। কারণ মাঝের অ্যাসাইনমেন্ট লাইনটি রান হওয়ার সাথে সাথে মেমরির সেই undefined মানটি আপডেট হয়ে একটি রিয়েল ফাংশন বডিতে রূপান্তরিত হয়ে যায়।

### anonymous function inside IIFE (Immediately Invoked Function Expression)

- যেই function এর কোন নাম নাই তাকে anonymous funcion বলে।
- IIFE (তাত্ক্ষণিক রান হওয়া ফাংশন) কে একবারের বেশি সরাসরি call করা যায় না।
- Creation Phase-এ গ্লোবাল মেমরিতে এদের নিজস্ব নামের কোনো অ্যালোকেশন বা হোইস্টিং হয় না।

```js
(function () {
  console.log("I am anonymous function inside an IIFE");
})();
```

### arrow function

function কিওয়ার্ডটি বাদ দিয়ে একটি তীর চিহ্ন (=>) ব্যবহার করা হয়, তাই এর নাম অ্যারো ফাংশন। <br>
Arrow function ও এক ধরনের function expression

```js
var print = () => {
  console.log("I am arrow function");
};
```

### parameters vs arguments

```js
var output = (text) => {
  // (text) is parameters that is recevied the arguments
  console.log(text);
};
output("Hello"); // arguments that is passed to parameter
```

### first order, higher order and call-back function

জাভাস্ক্রিপ্টে ফাংশনকে যখন সাধারণ ভ্যারিয়েবলের মতো অন্য ফাংশনের আর্গুমেন্ট হিসেবে পাস করা যায় বা কোনো ফাংশন থেকে রিটার্ন করা যায়, তখন সেই ক্ষমতা বা বৈশিষ্ট্যকে First-Class Function বলে।

Higher-Order Function: যে ফাংশনটি অন্য ফাংশনকে প্যারামিটার হিসেবে গ্রহণ করে বা রিটার্ন করে (যেমন আপনার কোডের: call ফাংশন)।

Callback Function: যে ফাংশনটিকে অন্য ফাংশনের ভেতরে আর্গুমেন্ট হিসেবে পাঠানো হয় (যেমন আপনার কোডের: sum ফাংশন)।

```js
function call(add) {
  // higher order function
  add(4, 5);
}

function sum(a, b) {
  // callback
  console.log(a + b);
}
call(sum); // first class function
```
