/**
 * Global scope and local scope
 */

var a = 10; // global variable
// global function
function outer() {
  // local function | function scope
  function inner() {
    var c = 40; // local variable
  }
  console.log(a); // can accesible cause a is define in global scope
  var b = 20; // local variable | function scope
  console.log(b); // can access cause it's under it's own local scope
}
console.log(b); // can't access it cause b is note define in global scope

/**
 * Block Scope
 */

{
  let name = "Arafat";
  const lName = "Hossain";
  var fullName = "Roki";
}
console.log(fullName); // can access
console.log(name); // not define in gloal

/**
 * shadwing
 */

var num = 43
function ac(){
    var num = 300
    console.log(num)
}
ac()