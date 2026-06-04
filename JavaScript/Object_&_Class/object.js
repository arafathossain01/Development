/* 

object একটা variable যার ভিতরে value এবং function দুইটাই রাখা যায়। 
key:value --> properties
key:function() --> methods.

*/

const user1 = {
  firstName: "Arafat",
  lastName: "Hossain",
  print: () => {
    console.log(this.firstName, " ", this.lastName);
  },
};
console.log(user1.firstName); // Arafat
user1.print(); // undefined undefined
/*
 Arrow function-এর নিজস্ব কোনো this থাকে না। সে তার বাইরের environment (lexical scope) থেকে this-কে ধার করে। এখানে this অবজেক্ট দুটিকে না বুঝিয়ে, গ্লোবাল Window অবজেক্টকে বোঝাবে। ফলে আউটপুট undefined undefined আসবে।
 */

const user2 = {
  firstName: "John",
  lastName: "Doe",
  age: 34,
  print() {
    console.log(this.firstName, " ", this.lastName);
  },
};
console.log(user2["firstName"]); // john
user2.print(); // john Doe

/**
 * firstName -> propertie
 * "John" -> value
 * কোনো অবজেক্টের ভেতরের মেথড থেকে যখন this-কে কল করা হয়, তখন this মানে হলো সেই অবজেক্ট নিজেই।
 */

user2.firstName = "Ikbal"; // update the property
user2.country = "Bangladesh"; // addad new property

console.log(user2.firstName);
console.log(user2.firstName);

delete user2.age; // delete করার পর property access করতে গেলে undefined দেখাবে।
console.log(user2);

let check = "college" in user2; // check college is exist or not in obect user2, if exist then return true other wise false.
console.log(check);

// loop for object
const person = {
  name: "John",
  age: 29,
  city: "New York",
};

let text = " ";
for (let x in person) {
  text += person[x] + " ";
}
console.log(text);

const arr = Object.values(person); // convert obj to array
console.log(arr);

const string = JSON.stringify(person); // obj to stringifyObj
console.log(string);

/*
    Object constructor

    একই রকমের অনেকগুলো অবজেক্ট (Object) সহজে এবং বারবার তৈরি করার জন্য Object Constructor বা কনস্ট্রাক্টর ফাংশন ব্যবহার করা হয়
 */

function Human(fName, lName, age) {
  this.firstName = fName;
  this.lastName = lName;
  this.age = age;

  this.fullName = function () {
    return this.firstName + " " + this.lastName;
  };
}
/*
কনস্ট্রাক্টর ফাংশনের ভেতর this হলো একটি খালি পাত্র (Placeholder)। new কিওয়ার্ড দিয়ে যখনই যার নামে অবজেক্ট বানানো হবে, this সাথে সাথে সেই অবজেক্টের রূপ ধারণ করবে।
 */

const human1 = new Human("Arafat", "Hossain", 23);
console.log(human1);

const human2 = new Human("John", "Doe", 29);
console.log(JSON.stringify(human2));

console.log("Full Name: " + human2.fullName());

human1.nationality = "Bangali"; // addad new property at human1 object
console.log(human1);

Human.prototype.nationality = "English"; // addad new property at Human constructor
console.log(human2.nationality);
