/*
যখন আমাদের একই রকমের অনেকগুলো অবজেক্ট তৈরি করতে হয়, তখন বারবার হাতে কোড না লিখে আমরা একটা class বানিয়ে নিই। class হচ্ছে একটা template.
The constructor method is called automatically when a new object is created.
*/
class Car {
  constructor(name, year, company) {
    this.name = name;
    this.year = year;
    this.company = company;
  }
  print() {
    console.log(`Car Name: ${this.name}
Model Year: ${this.year}
Company Name: ${this.company}
Family Name: ${this.family}`);
  }
}

// extend  use করে Car class কে ব্যাবহার করা হচ্ছে | Car class কে inherit করা হচ্ছে। 
class Premio extends Car {
  constructor(name, year, company, family) {
    super(name, year, company); // parent class কে refer করে। | প্যারেন্ট ক্লাসের কনস্ট্রাক্টরকে রান করায় | প্যারেন্ট ক্লাসের কাছে ডাটা বা আর্গুমেন্ট পাস করে
    this.family = family;
  }
}

class Allion extends Car {
  constructor(name, year, company, family) {
    super(name, year, company);
    this.family = family;
  }
}

const car1 = new Premio("F Premio", 2002, "Toyota", "Corona");
const car2 = new Allion("Allion", 2015, "Toyota", "Carina");

car1.print();
car2.print();
