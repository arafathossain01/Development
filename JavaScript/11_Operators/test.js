/* 
Arithmetic Operators

1. +	 Addition
2. -	  Subtraction
3. *	 Multiplication
4. **  Exponentiation (ES2016)
5. /	 Division
6. %    Modulus (Remainder)
7. ++  Increment
8. --	 Decrement

*/

const num1 = 20;
const num2 = 11;
const rem = num1 % num2;
console.log(rem);

const x = 5;
const z = x ** 2; //25
console.log(z);

/*
Assignment Operators

1. =	            x = y	        x = y	                x = 5
2. +=	           x += y	     x = x + y	           x = 15
3. -=	            x -= y	       x = x - y	          x = 5
4. *=      	       x *= y	     x = x * y	           x = 50
5. **=	          x **= y	  x = x ** y	       x = 100000
6. /=	            x /= y	      x = x / y              x = 2
7. %=	           x %= y	   x = x % y	         x = 0
8. :	              x: 45	        size.x = 45	        x = 45

*/

/*
Logical Operators

&& (and) --> return true if both true
|| (or) --> return true if any one is true
! (not) --> convert the boolean value 
*/

let a = true;
let b = false;
console.log(a && b); // false
console.log(a || b); // true
console.log(!b); // true

/*
Comparison Operators

1. ==         equal to
2. ===       equal value and equal type
3. !=          not equal
4. !==        not equal value or not equal type
5. >            greater than
6. <            less than
7. >=         greater than or equal to
8. <=         less than or equal to
*/
