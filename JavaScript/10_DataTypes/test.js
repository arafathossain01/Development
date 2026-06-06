/**
 * JavaScript has 2 types of DataTypes:
 * 1. Primitive             2. Object
 * 1.1 Boolean                  2.1 object
 * 1.2 null                         2.2 array
 * 1.3 undefined             2.3 function
 * 1.4 Number                    2.4 date
 * 1.5 BigInt                    2.5 regexp
 * 1.6 String                    2.6 set
 * 1.7 Symbol                    2.7 map
 */

const number = 29;
const boolean = true;
const string = `Hello SWE`;
const bigint = 9007199254740992n;

// Data type conversion
const num = 20; // number
console.log(num.toString()); // number to string
const number1 = 12345678;
console.log(number1.toLocaleString("de-DE")); // using for currency formatting

const str = `my age is ` + 20; // string
const str2 = 30 + ` is a number.`; //string
console.log(str);

const id = Symbol(); // every sumbol() has an unique value
console.log(id);

const num2 = "10" * 1; //number
console.log(typeof num2); // number

// string to number method
console.log(parseInt("101", 2)); // return whole number
console.log(parseFloat("2.3")); // return float number

const strNum = "42";
console.log(+strNum); // আউটপুট: 42 (Number) - খুব সংক্ষিপ্ত ও ফাস্ট

console.log(typeof NaN); // number

let user = null;
console.log(typeof user); // আউটপুট: "object" (এটিই সেই বিখ্যাত বাগ!)

/**
 * JS falsy value
 * 1. false
 * 2. 0
 * 3. ""
 * 4. null
 * 5. undefined
 * 6. NaN
 */

//falsy check
console.log(Boolean(NaN));

console.log(typeof { name: "SWE" }); // "object"
console.log(typeof [1, 2, 3]); // "object" (জাভাস্ক্রিপ্টে অ্যারেও এক ধরণের স্পেশাল অবজেক্ট!)
console.log(typeof function () {}); // "function"
