### package install

- npm i -g create-react-app [installed globally]
- npm install -g create-react-app [installed globally]
- npm i create-react-app [install locally]

### React setup

- npm i -g create-react-app
- create-react-app projectname [create a react application]
- npm start

### Create react project using vite

- npm create vite@latest
- select the project name
- select the framework
- select the language
- 


### way of clean

1. public
   - index.html
2. src
   - index.js
   - app.js

### React কে কেন Single Page Application বলা হয়?

রিয়্যাক্ট-কে Single Page Application (SPA) বলা হয় কারণ এই আর্কিটেকচারে পুরো অ্যাপ্লিকেশনের জন্য কেবল একটিমাত্র HTML ডকুমেন্ট (index.html) ব্রাউজারে লোড হয়।

যখন কোনো ইউজার এক পেজ থেকে অন্য পেজে নেভিগেট করেন, তখন ব্রাউজার সার্ভারে পুনরায় কোনো রিকোয়েস্ট পাঠায় না এবং পুরো পেজটি রিলোড হয় না। বরং রিয়্যাক্ট ক্লায়েন্ট-সাইড রাউটিং এবং জাভাস্ক্রিপ্টের সাহায্যে ওই একক ডকুমেন্টের রুট (#root) ডিভের ভেতরের কম্পোনেন্টগুলোকে ডায়নামিকালি আপডেট বা সোয়াপ (Swap) করে। এর ফলে ইউজার কোনো রিফ্রেশ ছাড়াই অত্যন্ত দ্রুত ও মসৃণভাবে ভিন্ন ভিন্ন ইন্টারফেস দেখতে পান। 
