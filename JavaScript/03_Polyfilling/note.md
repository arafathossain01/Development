### Compiling

Compiling হলো এমন একটি প্রক্রিয়া যেখানে একটি ভাষার code কে অন্য একটি language বা lower-level code এ রূপান্তর করা হয় যাতে computer সহজে execute করতে পারে।

JavaScript এর ক্ষেত্রে modern JavaScript code কে, older JavaScript version এ convert করা হতে পারে অথবা machine code এ optimize করা হতে পারে।

```js
const add = (a, b) => a + b; // modern js
```

```js
var add = function (a, b) {
  return a + b;
}; // older js
```

**babel, compile করে modern js কে older js এ রূপান্তরিত করে। এখানে code structure পরিবর্তন হয়েছে যাতে পুরোনো browser ও বুঝতে পারে।**

### Polyfilling

Polyfill হলো এমন code যা browser এ না থাকা কোনো নতুন feature কে manually add করে দেয়। যদি কোনো browser নতুন JavaScript feature support না করে, তাহলে Polyfill সেই feature emulate করে।

Polyfilling এর কাজ:

- missing feature add করা
- browser compatibility দেওয়া
- modern API emulate করা

ধরো পুরোনো browser এ Array.includes() নেই।

- **Array.includes()** polyfill

```js
const arr = [1, 2, 3];

arr.includes(2);
```

After polyfill

```js
if (!Array.prototype.includes) {
  Array.prototype.includes = function (value) {
    return this.indexOf(value) !== -1;
  };
}
```

- **map** polyfilling

```js
const arr = [1, 2, 3, 4];
const newArr = arr.map((elem) => {
  return elem * 2;
});
console.log(newArr);
```

After polyfilling

```js
Array.prototype.pMap = function (callback) {
  let res = [];
  for (let i = 0; i < this.length; i++) {
    res.push(callback(this[i], i, this));
  }
  return res;
};
```
