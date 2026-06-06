## JS Short-circuiting

Short-circuiting হলো JavaScript-এর এমন একটি behavior, যেখানে logical operator (&&, ||, ??) ব্যবহার করলে result নিশ্চিত হয়ে গেলে বাকি expression আর check করা হয় না।

1. AND (&&)

- প্রথম falsy value পেলেই evaluation stop হয়ে যায়

```js
true && "js"; // "js"
false && "node"; // false
```

- সবগুলো true হলে → last value return করে
- কোনো falsy পেলেই → সেটাই return করে

2. OR (||)

- প্রথম truthy value পেলেই evaluation stop হয়

```js
"" || "Default"; // "Default"
true || "JS"; // true
```

- কোনো truthy value পেলেই → সেটাই return
- সব falsy হলে → last value return

3. Nullish Coalescing (??)

- শুধু null বা undefined হলে right side return করে

```js
null ?? "Guest"; // "Guest"
"" ?? "Guest"; // ""
0 ?? 10; // 0
```

- null/undefined → right value
- অন্য কিছু (0, "", false) → left value
