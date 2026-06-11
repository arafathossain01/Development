// js number 64-bit floating point

let x = 0.2 + 0.1;
console.log(x); //0.30000000000000004
/**
 * এই সমস্যা টা হওয়ার কারণ হলঃ JavaScript number binary format এ store করে। 0.1 এর exact binary value তে convert না করতে পারায় সে একটা আনুমানিক নাম্বার ধরে নেয় এবং এই ভুল টা হয়। 0.1 আর 0.2 exact binary-তে represent হয় না।
 */

