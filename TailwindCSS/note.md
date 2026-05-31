## install using CLI

<p><code>npm install tailwindcss @tailwindcss/cli</code></p>
create a src folder and create a input.css file inside the src folder.<br>
create a dist folder and create a output.css file inside the output folder. <br><br>

<p><code>npx @tailwindcss/cli -i ./src/input.css -o ./dist/output.css --watch</code></p>
in our project create a index.html file.


## how to write custom css in tailwind
bg-[custom here]

## using theme
in input.css file we write it
```js
@theme{
    --primary-color: custom_css / tailwind_css;
    --radius-btn: 1rem;
}
```
## using layer
layer use for good readability
```js
@layer components{
    .class_name{
        border-radius: var(--radius-btn);
        @apply inline-flex
    }
    .class_name{
        @apply tailwind-css-write-here;
    }
    .card:hover{
        @apply write-css;
    }

}
```