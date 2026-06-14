<h1 align="center">JavaScript Array </h1>

Array হচ্ছে একটা **object type** data collection করার store. Array এর কিছু characteristics আছে :

1.  Array এর ভিতরে থাকা value গুলাকে element বলে।
2.  Array এর index order আছে
3.  Array এর index শুরু হয় 0 থেকে।
4.  Array তে সাইজ ছোট বড় করা যায় ভ্যালু ঢুকিয়ে বা বের করইয়ে।
5.  Array তে আলাদা আলাদা ধরনের data রাখা যায়।

```js
const arr = ["Sadnima", "Kaniz", "Nowshin"]; // array creation
let name = arr[0]; // access the first element of array
```

array এর ভিতরে loop ঘুরানোঃ

```js
const cars = ["Volvo", "Toyota", "Hino", "Lexus", "BMW"];
let lnth = cars.length;
for (let i = 0; i < lnth; i++) {
  console.log(cars[i]);
}

cars[cars.length - 1]; // last element
cars.length; // Array size
```

### Array Method

- pop(), push()

```js
const arr = ["mango", "banana", "apple", "orange"];

arr.pop();
console.log(arr);

arr.push("kiwi", "banana");
console.log(arr);
```

pop() মেথড array-এর শেষ (last index) থেকে একটি element remove করে। এতে array-এর length ১ কমে যায়।

push() মেথড array-এর শেষে এক বা একাধিক নতুন element যোগ করে।
এতে array-এর length বেড়ে যায়।

- shift() unshift()

```js
const arr = ["mango", "banana", "apple", "orange"];

arr.shift();
console.log(arr);

arr.unshift("kiwi", "banana");
console.log(arr);
```

shift() মেথড array-এর শুরু (index 0) থেকে প্রথম elementটা remove করে ফেলে। এটা হওয়ার পর বাকি সব element এক ধাপ বামে সরে আসে, মানে তাদের index এক করে কমে যায়।

unshift() মেথড array-এর শুরুতে এক বা একাধিক নতুন element যোগ করে।
এর ফলে আগের সব element এক বা একাধিক ধাপ ডানে সরে যায়, মানে তাদের index বেড়ে যায়।

- map()

```js
const arr = [10, 20, 30, 40, 50, 60, 70, 80, 90];
arr.map(find);

function find(value, index, array) {
  console.log(value * 2, index);
}
```

map() মেথড array-এর প্রতিটি element-এর উপর একটি function চালায়
এবং সেই function-এর return করা মান দিয়ে একটি নতুন array তৈরি করে। map() যে ৩টি argument নেয়:
<br>value → বর্তমান element-এর মান
<br>index → সেই element-এর index নম্বর
<br>array → পুরো array নিজেই

- filter()

```js
const arr = [10, 20, 30, 40, 50, 60, 70, 80, 90];
arr.filter(find);

function find(value, index, array) {
  console.log(value > 18, index);
}
```

filter() মেথড array-এর প্রতিটি element-এর উপর একটি condition চালায়। যে element-এর জন্য condition true হয়, শুধু সেই element গুলোকে নিয়ে একটি নতুন array তৈরি করে। filter() যে ৩টি argument নেয়:
<br>value → বর্তমান element-এর মান
<br>index → সেই element-এর index নম্বর
<br>array → পুরো array নিজেই

- find()

```js
const arr = [10, 20, 30, 40, 50];

const result = arr.find((value) => value > 25);
console.log(result); // 30
```

find() মেথড array-এর প্রতিটি element-এর উপর condition চালায় এবং যেই প্রথম element-এর জন্য condition true হয়, সেই element টাকেই return করে। condition match না করলে undefined return করে। find() যে ৩টি argument নেয়:
<br>value → বর্তমান element-এর মান
<br>index → সেই element-এর index নম্বর
<br>array → পুরো array নিজেই

- forEach()

```js
const fruits = ["apple", "banana", "mango"];

fruits.forEach((fruit) => console.log(fruit));

const arr = [
  "rafi",
  "miraz",
  "arafat",
  "shamim",
  "rafida",
  "jenin",
  "puja",
  "pranto",
  "najnin",
];

arr.forEach(find);

function find(value, index, array) {
  console.log(value, index, array);
}
```

forEach() হলো array-এর জন্য ব্যবহৃত একটি loop method। এটা array-এর প্রতিটি element-এর জন্য একটি function চালায়। forEach() কিছু return করে না। শুধু কাজ (print, update, calculation ইত্যাদি) করার জন্য ব্যবহার হয়। প্রতিটি element-এর জন্য function run করে। forEach() যে ৩টি argument নেয়:
<br>value → বর্তমান element-এর মান
<br>index → সেই element-এর index নম্বর
<br>array → পুরো array নিজেই

- includes()

```js
const fruits = ["apple", "banana", "mango"];

console.log(fruits.includes("banana")); // true
console.log(fruits.includes("orange")); // false
```

includes() হলো JavaScript-এর একটি array (এবং string) method,
যেটা দিয়ে কোনো নির্দিষ্ট value array-এর ভেতরে আছে কিনা তা check করা হয়। value থাকলে → true, value না থাকলে → false return করে. strict comparison (===) ব্যবহার করে

- reduce()

```js
const users = [
  { name: "A", age: 20 },
  { name: "B", age: 25 },
];

const totalAge = users.reduce((sum, u) => sum + u.age, 0);

console.log(totalAge); // 45
```

reduce() হলো JavaScript-এর একটি method যা array-এর অনেকগুলো value নিয়ে একটা single value তৈরি করে। এটি সাধারণত sum, multiplication, object aggregation, string concatenation ইত্যাদির জন্য ব্যবহার হয়। reduce() যে 4টি argument নেয়:
<br>accumulator → আগের সব result বা aggregated value
<br>currentValue → বর্তমান element
<br>index → বর্তমান element-এর index (optional)
<br>array → পুরো array (optional)

- some()

```js
const numbers = [45, 4, 9, 16, 25];

let allOver18 = numbers.some(function (value) {
  return value > 18;
});

console.log(allOver18); // true
```

some() হলো JavaScript-এর একটি array method,
যা চেক করে array-এর অন্তত একটি element কোনো condition পূরণ করছে কিনা।

যদি কমপক্ষে একটি element condition পূরণ করে → true

যদি কোনো element না মিলে → false

- every()

```js
const numbers = [45, 4, 9, 16, 25];

let allOver18 = numbers.every(function (value) {
  return value > 18;
});

console.log(allOver18); // false
```

every() হলো JavaScript-এর একটি array method,
যা চেক করে array-এর সব element কোনো condition পূরণ করছে কিনা।

যদি সব element condition পূরণ করে → true

যদি একটিও element না মিলে → false

- join()

```js
const arr = ["mango", "banana", "apple", "orange"];
let store = arr.join(" * ");
console.log(store);
```

join() হলো JavaScript-এর একটি array method, যা array-এর সব element কে একটি string-এ পরিণত করে এবং তাদের মধ্যে আপনি যেটা দেন তা separator হিসেবে যোগ করে।

যদি separator না দেওয়া হয় → default হচ্ছে ,

মূল array পরিবর্তন হয় না

- indexof()

```js
const fruits = ["apple", "banana", "mango", "banana"];

console.log(fruits.indexOf("banana", 2)); // 3
console.log(fruits.indexOf("orange")); // -1
```

indexOf() হলো JavaScript-এর একটি array (এবং string) method,
যা কোনো element array-এ কোথায় আছে (index) তা খুঁজে বের করে।

যদি element পাওয়া যায় → index number return করে

যদি element না থাকে → -1 return করে

- concat()

```js
const myGirls = ["Cecilie", "Lone"];
const myBoys = ["Emil", "Tobias", "Linus"];

const myChildren = myGirls.concat(myBoys);
console.log(myChildren);
```

concat() method is used to add two or more array.

- slice() ,

```js
const arr = ["apple", "banana", "orange", "jackfruit", "mango", "dragon-fruit"];

const separate = arr.slice(1, 4);

console.log(separate);
```

slice() হলো array থেকে নির্দিষ্ট অংশ কপি করে নতুন array তৈরি করার জন্য। মূল array change হয় না. start index থেকে শুরু করে end index আগে পর্যন্ত (exclude end) copy করে

start → কোন index থেকে শুরু করবে (included)

end → কোন index পর্যন্ত যাবে (excluded)

- sort()

```js
const num = [40, 54, 34, 10, 100];

let x = num.sort((a, b) => a - b); //asecending
num.sort((a, b) => b - a); //descending
console.log(x);
```

ডিফল্টভাবে JavaScript-এর sort() method সংখ্যাগুলোকে string হিসেবে ধরে sort করে। তাই number sort করলে অনেক সময় ভুল order পাওয়া যায়।
