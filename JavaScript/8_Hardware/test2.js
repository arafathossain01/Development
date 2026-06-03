console.log("start");


setTimeout(function print1(){ // setTimeout is first class function and print is callback function
    console.log("Yesss1");
}, 0); // ekta function, ekta time

setTimeout(function print2(){ // setTimeout is first class function and print is callback function
    console.log("Yesss2");
}, 4000); // ekta function, ekta time

console.log("End");
