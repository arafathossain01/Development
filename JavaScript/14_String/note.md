### string methods

- length

```js
let text = `Arafat Hossain`;
console.log(text.length);
```

length প্রোপার্টি একটি স্ট্রিং-এর দৈর্ঘ্য (অক্ষরের সংখ্যা) রিটার্ন করে।

---

string থেকে char এর position বের করে নিয়া আসার জন্য এই method গুলা ব্যাবহার করা হয়।

- at()

```js
const name = `myname`;
let letter = name.at(2);
console.log(letter);
```

at() method negative index নিয়ে কাজ করতে পারে। negative index use করলে last char থেকে counting start হবে। index counting শুরু হয় 0 থেকে।

- []

```js
const name = `myname`;
console.log(name[2]);
```

specific position থেকে char বের করে নিয়ে আসে। negative index নিয়ে কাজ করতে পারে নাহ। negative index number দিলে undefined দেখাবে।

- charAt()

```js
const name = `myname`;
let letter = name.charAt(2);
console.log(letter);
```

specific position থেকে char বের করে নিয়ে আসে। negative index নিয়ে কাজ করতে পারে নাহ।

- codePointAt()

```js
let text = "HELLO WORLD";
let code = text.codePointAt(0);
```

string-এর কোনো character-এর আসল Unicode code point বের করা

- fromCharCode()

```js
console.log(String.fromCharCode(65)); // A
```

code থেকে character বানায়

---

- concat()

```js
let text1 = "Hello";
let text2 = "World";
let text3 = text1.concat(" ", text2);
```

এক বা একাধিক string কে জোরা লাগায়।

---

string কে break করার জন্য এই method গুলা ব্যাবহার করা হয়।

- slice()

```js
let text = "Apple, Banana, Kiwi";
let part = text.slice(7, 13); // Banana
```

slice(start, end). start কোন জায়গা থেকে শুরু করবে। end কোনটার আগে শেষ করবে। js count position from 0. negative position number ও count করতে পারে এই method.

- substring()

```js
let str = "Apple, Banana, Kiwi";
let part = str.substring(7, 13); // Banana
```

substring() কোনো string-এর একটা অংশ বের করে আনে. negative মানে 0 ধরে নেয়।

---

লিখা ছোট বড় করার জন্য এইগুলা ব্যাবহার করা হয়।

- toUpperCase()

```js
let text1 = "Hello World!";
let text2 = text1.toUpperCase(); // HELLO WORLD
```

- toLowerCase()

```js
let text1 = "Hello World!";
let text2 = text1.toLowerCase(); // hello world
```

---

একটা string এর আগে পরে থেকে white space remove করার জন্য ৩ টা মেথড আছে।

- trim()

```js
let text1 = "      Hello World!      ";
let text2 = text1.trim();
```

সামনে পিছনে দুই দিক থেকেই whitespace বাদ করে দিবে।

- trimStrat()

```js
let text1 = "      Hello World!      ";
let text2 = text1.trimStart();
```

শুধু সামনে থেকে ফাঁকা জায়গা মুছে ফেলবে।

- trimEnd()

```js
let text1 = "      Hello World!      ";
let text2 = text1.trimEnd();
```

শুধু পিছন থেকে ফাঁকা জায়গা মুছে ফেলবে।

---

- matchAll()

```js
let str = "Hello World";

console.log(str.includes("World")); // true
console.log(str.includes("world")); // false → case-sensitive
console.log(str.includes("Hello", 1)); // false → 1 index থেকে খুঁজেছে
```

কোনো string বা array এর মধ্যে নির্দিষ্ট value আছে কিনা চেক করে. Boolean (true/false) return করে. Original string বা array change হয় না।

---

- split()

```js
let str = "apple,banana,mango";
let arr = str.split(",");
console.log(arr); // ["apple", "banana", "mango"]
```

split() কোনো string কে parts-এ ভাঙে separator এর ভিত্তিতে ভাগ করে array return করে Original string change হয় না

---

- replace()

```js
let str = "Hello World";
let newStr = str.replace("World", "JS");

console.log(newStr); // "Hello JS"
console.log(str); // "Hello World" (original unchanged)

let str = "apple apple apple";
str.replace(/apple/g, "orange"); // "orange orange orange"
```

replace() কোনো string-এর মধ্যে text খুঁজে প্রথমে match করে নতুন text এর সাথে, পরে নতুন text দিয়ে replace করে new string return করে।

---

search করতে যেইগুলা দরকার হয়।

- indexOf()

```js
let text = "Hello World";
console.log(text.indexOf("World")); // 6
```

indexof() মেথড position return করে, মিলে যাওয়া ১ম char এর position. index খুঁজে না পেলে return -1 করে।

- lastIndexOf()

```js
let text = "apple banana apple";
console.log(text.lastIndexOf("apple")); // 13
```

শেষ এর থেকে match করবে এবং index return করবে।

- includes()

```js
let text = "Hello World";
console.log(text.includes("World")); // true
```

substring আছে কি না check করে থাকলে true return করে না থাকলে false return করে।

- search()

```js
let text = "Hello World";
console.log(text.search("World")); // 6
```

indexOf এর মতো কিন্তু regex support করে

---

- repeat()

```js
let text = "Hi ";
console.log(text.repeat(3)); // Hi Hi Hi
```

একটা element কে number পর্যন্ত বার বার return করবে।

- padStart()

```js
let num = "5";
console.log(num.padStart(3, "0")); // 005
```

padStart() string-এর বাম দিকে fill / যোগ করে যতক্ষণ না target length হয়।

- padEnd()

```js
let num = "5";
console.log(num.padEnd(3, "0")); // 500
```

padEnd() string-এর ডান দিকে fill / যোগ করে যতক্ষণ না target length হয়।

---

checking এর ক্ষেত্রে দরকার

- startsWith()

```js
let text = "Hello World";
console.log(text.startsWith("Hello")); // true
```

প্রথম থেকে string match করাবে, মিললে true return করবে। নয়তো false.

- endsWith()

```js
let text = "Hello World";
console.log(text.endsWith("World")); // true
```

শেষ থেকে string match করাবে, মিললে true return করবে। নয়তো false.

---

- replace()

```js
let str = "apple apple apple";
console.log(str.replace("apple", "orange"));
// orange apple apple
```

শুধু first match replace করে

```js
let str = "apple apple apple";
console.log(str.replaceAll("apple", "orange"));
// orange orange orange
```

সব match replace করে

---

- valueOf()

```js
let str = "Hello";
console.log(str.valueOf());
```

Returns the primitive value of a string or a string object
