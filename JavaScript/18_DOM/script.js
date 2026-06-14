// Selecting HTML element

const id = document.getElementById("demo"); // finding element using id
const tag = document.getElementsByTagName("p"); // finding element using tag name
/**
 * First, find the element with id="x",
 * then find all <span> tags inside it.
 */
const x = document.getElementById("x");
const tags = x.getElementsByTagName("span");

const child = document.getElementsByClassName("child"); // find all same class name element

const css = document.querySelector(".css-selector"); // find element using css selector

const selector = document.querySelectorAll(".child"); // node list

const y = document.querySelectorAll("p.para"); // node list

const img = document.getElementsByTagName("img");

// Changing HTML contecnt

x.innerHTML = "<h1>Hello World</h1>";

img[0].src = "./nature2.jpg";

child[1].innerHTML = "<h1>Date: </h1>" + Date();

// changing CSS

y[1].style.background = "blue"


function colorChange(){
    y[1].innerHTML = "<h1>After click i am here</h1>"
    y[1].style.text = "hidden"
}