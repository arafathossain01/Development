### Internal Memory

Go এর internal memory ৪ ভাগে বিভক্তঃ

1. code segment
2. data segment
3. stack
4. heap

এছাড়া Heap memory manage করার জন্য Go runtime এ **Garbage Collector (GC)** থাকে।

- **Code Segment** 

    Code Segment এ program এর compiled machine instructions থাকে।

    এখানে সাধারণত থাকে:

    - function code
    - program instructions
    - constant values (read-only)

- **Data Segment**

    Data Segment এ global এবং static data store হয়।

    এখানে সাধারণত থাকে:

    - global variable
    - package level variable

- **Stack**

    যখন Code Segment থেকে কোনো function execute হয়, তখন সেই function এর execution data Stack এ যায়।

    Stack এ থাকে:

    - local variable
    - function parameter
    - return address
    - temporary data

    প্রতিটি function call, Stack এ একটি নির্দিষ্ট জায়গা দখল করে, যাকে Stack Frame বলে। যখন function এর কাজ শেষ হয় তখন stack frame automatically remove হয়ে যায় এটাকে stack unwinding/pop বলা হয়।

- **Heap** <a href="../Functions/closure/main.go" style="color:red">Clouser</a>

    Heap এ dynamic memory allocation হয়।

    যে data runtime এ তৈরি হয় এবং function শেষ হওয়ার পরও দরকার হতে পারে, সেগুলো Heap এ যায়।

*Garbage Collector (GC)*<br>
Heap memory manually free করতে হয় না।

Go এর Garbage Collector:

unused heap memory detect করে<br>
automatically clean করে

<img src="./garbage_collector.webp" heigh="200px" width="200px" alt="garbage collector">
