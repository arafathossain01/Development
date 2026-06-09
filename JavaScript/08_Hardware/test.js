var x = 10;

function outer() {
  var a = 11;
  function inner() {
    var b = 12;
    function mostInner() {
      var c = 13;
      console.log(x);
      console.log(a);
      console.log(b);
      console.log(c);
    }
    mostInner();
  }
  inner();
}
outer();
var y = x;
console.log(y);

/*
    এখানে global এর জন্য কোন closure create করবে নাহ। বাকি সবার জন্য clouser তৈরি হয়ে কাজ করবে কারণ সে বার বার নিচে আসবে নাহ। 
*/
