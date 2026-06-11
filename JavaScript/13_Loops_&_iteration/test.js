// for loop
/**
 * যখন Count fixed (নির্দিষ্ট) থাকবে, তখনই আমরা for loop ব্যবহার করব
 */
for (let i = 1; i <= 10; i++) {
  console.log(i);
}

// while loop

/*let i = 1;
while (i<=10){
    console.log(i);
    i++;
}*/

// do.....while loop

/*
 *do while loop must be run at least once
 */
let j = 1;
do {
  console.log(j);
  j++;
} while (j <= 2);

// for....in loop | Object এর key (index/property name) iterate করার জন্য use হয়
const obj = {
  fName: "Arafat",
  lName: "Hossain",
  age: 23,
};

for (let key in obj) {
  console.log(key); // fName lName age
  console.log(obj[key]); // Arafat Hossain 23
}

// for.....of loop | Iterable (array, string, map, set) এর value পাওয়ার জন্য
const arr = [1, 2, 3, 4, 5, 6, 7];
for (let n of arr) {
  console.log(n);
}
