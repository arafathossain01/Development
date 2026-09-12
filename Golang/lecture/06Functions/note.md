- **Higher order function (first class function) VS first order function**
  - Higher order function following any one of three: <a href="./HOfunction/main.go" >HOfunction</a>

    i. parameter হিসাবে function যাবে<br>
    ii. function return করবে<br>
    iii. both

  - First order functions are simple, they work with:

    i. normal data/ value <br>
    ii. function parameter নেয় না <br>
    iii. function return করে নাহ।

- Callback function

  Higher Order Function এর ভিতরে parameter হিসেবে যে function পাঠানো হয়, সেটাকেই সাধারণভাবে Callback Function বলে।

- First class citizen

  যাদের variable এ assigen করা যায়, function এর argument হিসাবে পাঠানো যায়, function থেকে return করা যায় তাদের first class citizen বলে। যেহেতু go তে function কে ও return করতে পারি, variable এ assigne করতে পারি তাই higher order funtion কে first class function বলা হয়।

- Variadic function

- Closure - close over <a href="./closure/main.go">Colsure</a> <a href="../InternalMemory/note.md" style="color:red">[InternalMemory]</a>
- Defer function - last in first out
- Receiver function <a href="./receiver/main.go">Receiver</a> <a href="../Structure/main.go" style="color:red">[know first struct]</a>
