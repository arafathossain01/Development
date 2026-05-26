### var

```js
var a = 10;

if (true) {
  var p = 10;
  console.log(p);
}

console.log(a);
console.log(p); // can access though it's a block scoped variable
```

`var` একটা function scoped.`var` কে function এর বাইরে access করা যায় নাহ। কিন্তু block এর বাইরে একে ঠিকই access করা যায়, যার ফলে কিছু bug তৈরি হয় unexpected variable access তৈরি হয়। JavaScript এ code execute হওয়ার আগে Creation Phase এ`var` memory তে জায়গা পায় এবং initial value হিসেবে undefined assign হয়। এরপর Execution Phase এ actual value assign করা হয়। তাই `var` কে block এর বাইরে থেকে ও access করা যায়। এই access বন্ধ করার জন্য পরে 2015 সালে `let` & `const` কে নিয়ে আসা হয়।

### let & const

2015 সালে let কে নিয়ে আসা হয় যেন block এর ভিতরের কোন কিছু বাইরে access করা না যায়।

```js
console.log(a); //ReferenceError: Cannot access 'a' before initialization

let a = 10;

console.log(a);
```

`let` দিয়ে declare করা variable hoisting হয় কিন্তু global execution context এ হয় নাহ সেটা script নামে আলাদা একটা scope এ হয়, creation phase এ variable এর value undefined থাকে এবং যতক্ষণ না real value assign হয় ততক্ষণ variable access করা যায় নাহ তাই declare এর আগে access করা যাবে নাহ `let` variable কে। কারণ তখন `let` variable **Temporal Dead Zone** (Time) এ থাকে। Temporal Dead Zone হল যখন থেকে `let` দিয়ে declare করা variable creation phase এ undefined হয় তখন থেকে শুরু করে variable এ actual value assign হওয়ার আগ পর্যন্ত সময়টা হচ্ছে Temporal Dead Zone. তাই deaclare এর আগে variable access করা হলে ReferenceError দেয়।

```js
let p = 10;
if (true) {
  let a = 10;
  console.log(a);
}

console.log(a);
```

let p গ্লোবাল এক্সিকিউশন কন্টেক্সটের ক্রিয়েশন ফেজে স্ক্রিপ্ট স্কোপে জায়গা পায় এবং এক্সিকিউশন ফেজে তার মান ১০ হয়। পরবর্তীতে কোড রান হতে হতে যখন if ব্লকের ভেতরে প্রবেশ করে, তখন সাময়িকভাবে একটি নতুন Block Scope তৈরি হয় এবং তার ভেতরে let a হোইস্টিং বা তৈরি হয়। ব্লক থেকে কোড বের হয়ে গেলে এই ব্লক স্কোপটি মেমোরি থেকে ডিলিট হয়ে যায়, তাই ব্লকের বাইরে a-কে আর পাওয়া যায় না।

`const` এবং `let` এর কাজের ধরন প্রায় একই রকম—উভয়ই Block Scoped এবং উভয়েই হোইস্টিং হওয়ার পর Script Scope বা Block Scope-এ জমা হয়ে Temporal Dead Zone (TDZ) তৈরি করে। শুধু difference `const` দিয়ে ডিক্লেয়ার করলে তার মান জীবনেও পরিবর্তন বা Re-assign করা যায় না। আর তাকে ডিক্লেয়ার করার লাইনেই মান অ্যাসাইন করতে হয়, ফাঁকা রাখা যায় না।

**code করার সময় সব সময় `const` use করতে হবে যদি variable এর মান change করতে চাই তখন `let` use করতে হবে।**

- let, const
  - block scope / script scope
- var
  - global scope / local scope
