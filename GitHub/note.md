- গিট তার ডাটাবেজে কোন কিছু ফাইলের নাম ধরে সংরক্ষন না করে হ্যাশ কোড(24b9da6552252987aa493b52f8696cd6d3b00373) আকারে সংরক্ষন করে।

- **Modified (পরিবর্তিত):** আপনি কোড এডিটরে বসে একটা ফাইলের লাইন পরিবর্তন করলেন বা নতুন কিছু লিখলেন। ফাইলটি সেভ হয়েছে, কিন্তু Git এখনো এটাকে তার পরবর্তী স্ন্যাপশটের জন্য অফিসিয়ালি হিসেব করেনি। এটি থাকে আপনার Working Directory-তে।

- **Staged (চিহ্নিত):** আপনি git add কমান্ড চালিয়ে Git-কে বললেন, "আমি এই এই পরিবর্তনগুলো চূড়ান্ত করতে চাই, এগুলোকে পরের চালানের জন্য রেডি করো।" এটি জমা হয় Staging Area বা ইনডেক্সে।

- **Committed (সংরক্ষিত):** আপনি git commit করলেন এবং একটি মেসেজ দিলেন। ব্যস! আপনার কোডের ওই মুহূর্তের একটি স্থায়ী স্ন্যাপশট আপনার কম্পিউটারের লোকাল গিট ডাটাবেজে (.git ফোল্ডারে) চিরতরে সুরক্ষিত হয়ে গেল।

### গিটের কনফিগারেশন সাধারণত তিনটি আলাদা স্তরে সংরক্ষিত থাকে:

- System: পুরো কম্পিউটারের সব ইউজারের জন্য।

- Global: আপনার কম্পিউটারের বর্তমান ইউজারের সব প্রজেক্টের জন্য।

- Local: শুধুমাত্র আপনি এখন যে নির্দিষ্ট প্রজেক্ট ফোল্ডারে আছেন সেটির জন্য।

### Command line and their work

- git config --list --show-origin [show settings list and their path]
- git config --global user.name "Name" [set user name at git]
- git config --global user.email email@example.com [set email address at git]
- git help config [show the official document of git config]
- git add -h [show the help manual of git add]