<h1 align="center">JavaScript Number </h1>

JavaSCript এ এক ধরনের number ই আছে শুধু (double)। এটা দশমিক সহ লিখা যায় আবার ছাড়া ও লিখা যায়। JavaScript number হলো 64-bit floating point, এতে প্রায় 15–17 digit পর্যন্ত precision ঠিক থাকে।

```js
let num = 12;
let num1 = 12.12;
let num2 = 123e5; //12300000  large number
let num3 = 123e-5; // 0.00123  small number
```

```js
let y = 9999999999999999; // y will be 10000000000000000
```

JavaScript এ Number.MAX_SAFE_INTEGER = 9007199254740991 এটা হোল safe range এর বেশি নাম্বার দিলে সে ঠিক ভাবে ধরে রাখতে পারে নাহ, তাই 9999999999999999 এর কাছাকাছি value হিসাবে 10000000000000000 print করে।

```js
let sum = 0.1 + 0.2; //0.30000000000000004

// solve

let solveSum = (0.1 * 10 + 0.2 * 10) / 10;
console.log(solveSum); // 0.3
```

এই সমস্যা টা হওয়ার কারণ হলঃ JavaScript number binary format এ store করে। 0.1 এর exact binary value তে convert না করতে পারায় সে একটা আনুমানিক নাম্বার ধরে নেয় এবং এই ভুল টা হয়। 0.1 আর 0.2 exact binary-তে represent হয় না।

---

```js
let num = 10 + "10"; // 1010
```

string এবং number + operator দিয়া যোগ করলে concat হয়ে যায়। শুধু number + operator দিয়ে যোগ করলে সঠিক ফলাফল দেখাবে।

**Not a Number**

Not a Number যাকে NaN বলা হয়। JavaScript এর একটা reserved word. অনেক ভাবে একটা নাম্বার NaN হতে পারে। <br>
NaN এর টাইপ হল Number.

```js
let x = 100 / "shill"; //NaN
let x = 100 / "10"; // it is a number and return 10
let x = NaN + 10; // NaN
```

### Number Method

- toString()

```js
let num = 19;
num.toString();
console.log(num); // give it string
```

num.toString() শুধু string return করে, original num number হিসেবে থাকে।

- toFixed()

```js
let x = 9.3463;
let y = x.toFixed(2);
console.log(typeof y); // 9.35 and the type is string
```

the toFixed() method returns a string with a specified number of decimals.

- Number()

```js
Number("123"); // 123
Number("12.5"); // 12.5
Number("abc"); // NaN
```

The Number() method can be used to convert JavaScript variables to number

- parseInt()

```js
parseInt("42px"); // 42
parseInt("101", 2); // 5 (binary)
```

parseInt() Converts string to integer only, ignores trailing non-numeric characters.

- parseFloat()

```js
parseFloat("19.99$"); // 19.99
```

parseFloat() convert the string to decimal number.

- Number.isNaN()

```js
Number.isNaN(NaN); // true
Number.isNaN("abc"); // false
```

Strict NaN check (only true if value is actually NaN)

- Math.round()

```js
Math.round(4.6); // 5
```

rounds to nearest integers.

- Math.floor()

```js
Math.floor(4.9); // 4
```

Math.floor() round a number down.

- Math.ceil()

```js
Math.ceil(4.1); // 5
```

Upward round, negative number ক্ষেত্রে কম negative হয়।

- toLocalString()

```js
(1500).toLocaleString("en-IN", {
  style: "currency",
  currency: "INR",
});
// "₹1,500.00"
```

User-friendly format বানায়
comma, currency, country format অনুযায়ী

- Number.MAX_VALUE / Number.MIN_VALUE

```js
Number.MAX_VALUE;
// 1.7976931348623157e+308
Number.MIN_VALUE;
// 5e-324
```

Number.MAX_VALUE JavaScript-এ সবচেয়ে বড় possible number এর চেয়ে বড় হলে value হয় Infinity <br>
JavaScript-এ সবচেয়ে ছোট POSITIVE number এটা negative না 0 এর থেকে একটু বড়
