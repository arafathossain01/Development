### JavaScript কি?

জাভাস্ক্রিপ্ট একটি জনপ্রিয় ও শক্তিশালী প্রোগ্রামিং ভাষা, যা মূলত ওয়েবসাইটকে ইন্টারঅ্যাকটিভ ও ডাইনামিক করার জন্য ব্যবহার করা হয়। এর মাধ্যমে ওয়েবপেজে বিভিন্ন ধরনের কাজ করা যায়, যেমন বাটনে ক্লিক করলে কিছু দেখানো, পপআপ মেনু তৈরি করা, অ্যানিমেশন চালানো, ফর্ম ভ্যালিডেশন করা এবং ব্যবহারকারীর বিভিন্ন কার্যক্রমের প্রতিক্রিয়া দেওয়া।

জাভাস্ক্রিপ্ট ব্রাউজারের ভিতরে HTML ও CSS এর সাথে কাজ করে ওয়েবপেজের ডিজাইন ও কনটেন্ট পরিবর্তন করতে পারে। এছাড়াও Node.js এর মাধ্যমে জাভাস্ক্রিপ্ট সার্ভার সাইড বা ব্যাকএন্ড ডেভেলপমেন্টেও ব্যবহার করা হয়। এর সাহায্যে ডাটাবেজ পরিচালনা, API তৈরি, রিয়েল-টাইম অ্যাপ্লিকেশন তৈরি এবং সার্ভার নিয়ন্ত্রণ করা সম্ভব।

বর্তমানে ওয়েব ডেভেলপমেন্ট, মোবাইল অ্যাপ, গেম এবং বিভিন্ন আধুনিক সফটওয়্যার তৈরিতে জাভাস্ক্রিপ্ট ব্যাপকভাবে ব্যবহৃত হচ্ছে।

### Compiler & interpreter

Compiler এবং Interpreter উভয়ই Source Code কে Machine Code এ রূপান্তর করে এবং প্রোগ্রাম রান করতে সাহায্য করে, তবে তাদের কাজের পদ্ধতি ভিন্ন।

- Compiler

  Compiler একটি High-Level Language কে একবারে Low-Level Language বা Machine Code এ রূপান্তর করে। পুরো প্রোগ্রাম আগে compile করা হয়, তারপর executable file তৈরি হয় এবং প্রোগ্রাম রান করা যায়।

  Compiler সাধারণত:
  - পুরো code একবারে translate করে
  - execution speed বেশি হয়
  - memory বেশি ব্যবহার করে
  - error compile time এ দেখায়

  উদাহরণ:
  - C
  - C++
  - C#
  - Java

  Process:

  `source code => compile => machine code => output`

- interpreter

  Interpreter High-Level Language কে line by line translate ও execute করে। অর্থাৎ program রান হওয়ার সময়ই translation সম্পন্ন হয়।

  Interpreter সাধারণত:
  - line by line execute করে
  - memory কম ব্যবহার করে
  - execution তুলনামূলক ধীর হয়
  - debugging সহজ হয়

  উদাহরণ:
  - JavaScript
  - Python
  - PHP
  - Ruby
  - Perl

  Process:

  `source code => interpreter => output`

### JIT (Just In Time Compilation)

JavaScript Engine এ JIT (Just-In-Time) Compilation হলো এমন একটি প্রক্রিয়া যেখানে কোড রান হওয়ার সময়ই সেটাকে optimize করা হয় যাতে program আরও দ্রুত execute করতে পারে।

প্রথমে JavaScript কোড Interpreter এর মাধ্যমে line by line execute হতে থাকে। এই সময় একটি profiler বা monitor লক্ষ্য করে কোন statement বা function সবচেয়ে বেশি বার রান হচ্ছে। যেসব অংশ বারবার execute হয় সেগুলোকে প্রথমে “Warm Code” এবং পরে “Hot Code” হিসেবে চিহ্নিত করা হয়।

এরপর JIT Compiler সেই Hot Code গুলোকে Machine Code এ compile করে ফেলে। ফলে পরবর্তীতে একই code আবার run হলে Interpreter কে আর line ধরে ধরে execute করতে হয় না। এতে execution speed অনেক বেড়ে যায়।

অর্থাৎ JIT Compilation এর মূল উদ্দেশ্য হলো program এর সবচেয়ে বেশি ব্যবহৃত অংশগুলোকে শনাক্ত করে সেগুলোকে optimized machine code এ রূপান্তর করা, যাতে JavaScript দ্রুত এবং efficient ভাবে কাজ করতে পারে।
