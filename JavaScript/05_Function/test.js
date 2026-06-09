// function statement or function declaration
function print() {
  console.log("Hello World");
}

// function expression or named function expression
var log = function print() {
  console.log("This is function expression");
};

// anonymous function inside IIFE (Immediately Invoked Function Expression)
(function () {
  console.log("I am anonymous function inside an IIFE");
})();

// arrow function

var print = () => {
  console.log("I am arrow function");
};

// parameters vs arguments

var output = (text) => {
  // (text) is parameters that is recevied the arguments
  console.log(text);
};
output("Hello"); // arguments that is passed to parameter

function call(add) {
  // higher order function
  add(4, 5);
}

function sum(a, b) {
  // callback
  console.log(a + b);
}
call(sum); // first class function
