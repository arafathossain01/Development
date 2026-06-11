let str1 = "Hello Wolrd"; // double quotes
let str2 = "I am noob"; // single quotes

const age = 20;
let str3 = `my age is ${age}`; // string template best practice
console.log(str1);
console.log(str2);
console.log(str3);

let str4 = `Hi, i'am Arafat`;
console.log(str4);

// String Method

let txt = `AbCdEfGhIj`;

let length = txt.length; // length property  10
console.log("string length: " + length);

console.log("UpperCase: " + txt.toUpperCase()); // convert to uppercase
console.log("LowerCase: " + txt.toLowerCase()); // convert to lowercase

let text = `        Hello Developer     `;
console.log("Trim: " + text.trim()); // remove space both side
console.log("Start trim: " + text.trimStart()); // remove space from start
console.log("End trim: " + text.trimEnd()); // remove space from end

console.log(txt.at(-2)); // char এর position বের করে. at() method negative index নিয়ে কাজ করতে পারে।
console.log(txt.charAt(2)); // char এর position বের করে. charAt() method negative index নিয়ে কাজ করতে পারে না।
console.log(txt[3]); //specific position থেকে char বের করে নিয়ে আসে।

let string = `Hello Bangladesh Hello`;
console.log(string.indexOf(`Bangladesh`)); //6 | indexof() মেথড position return করে, মিলে যাওয়া ১ম char এর position. index খুঁজে না পেলে return -1 করে।
console.log(string.lastIndexOf( `Hello`)); // 17 | শেষ এর থেকে match করবে এবং index return করবে।
console.log(string.includes("Hello")); // true | substring আছে কি না check করে থাকলে true return করে না থাকলে false return করে।
console.log(string.startsWith("Hello")); //প্রথম থেকে string match করাবে, মিললে true return করবে। নয়তো false.
console.log(string.endsWith("Bangladesh"));  // শেষ থেকে string match করাবে, মিললে true return করবে। নয়তো false.

console.log(string.slice(6,16)); //slice(start, end). start কোন জায়গা থেকে শুরু করবে। end কোনটার আগে শেষ করবে। js count position from 0. negative position number ও count করতে পারে এই method.
console.log(string.substring(6,16)); // substring() কোনো string-এর একটা অংশ বের করে আনে. negative মানে 0 ধরে নেয়।

let str = "apple apple apple";
console.log(str.replace("apple", "orange")); // শুধু first match replace করে
console.log(str.replaceAll("apple", "orange")); // সব match replace করে
console.log(str.split(",")) // split() কোনো string কে parts-এ ভাঙে separator এর ভিত্তিতে ভাগ করে array return করে Original string change হয় না
console.log(str.concat(" ",str1)); // এক বা একাধিক string কে জোরা লাগায়।


