// JavaScript has two conditional statements if...else and switch

/*
const number = 65;

if (number <= 100 && number >= 80) {
    console.log(`Your mark is: ${number}. You got A+`);
} 
else if (number < 80 && number >= 70) {
    console.log(`Your mark is: ${number}. You got A`);
} 
else if (number < 70 && number >= 60) {
    console.log(`Your mark is: ${number}. You got A-`); 
} 
else if (number < 60 && number >= 50) {
    console.log(`Your mark is: ${number}. You got B`);
} 
else if (number < 50 && number >= 33) {
    console.log(`Your mark is: ${number}. You got C`);
} 
else if (number < 33 && number >= 0) {
    console.log(`Your mark is: ${number}. You failed!`);
} 
else {
    console.log("Invalid number! Please enter a number between 0 and 100.");
}
*/

const number = 65;

switch (true) {
  case number <= 100 && number >= 80:
    console.log(`Your mark is: ${number}. You got A+`);
    break;

  case number < 80 && number >= 70:
    console.log(`Your mark is: ${number}. You got A`);
    break;

  case number < 70 && number >= 60:
    console.log(`Your mark is: ${number}. You got A-`);
    break;

  case number < 60 && number >= 50:
    console.log(`Your mark is: ${number}. You got B`);
    break;

  case number < 50 && number >= 33:
    console.log(`Your mark is: ${number}. You got C`);
    break;

  case number < 33 && number >= 0:
    console.log(`Your mark is: ${number}. You failed!`);
    break;

  default:
    console.log("Invalid number! Please enter a number between 0 and 100.");
}

// break statement লুপ বা switch-এর ভিতর থেকে একদম বের করে দেয়।
// continue statement লুপের current iteration skip করে, এবং পরবর্তী iteration চালু করে।


// 
let age = 20;
let result = (age < 18) ? console.log("You are not a voter") : console.log("You are a voter");