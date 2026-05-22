### Differnce between var and :=

| Feature | `var` ✅             | `:=` ❌ (limited)     |
| ------- | ------------------- | -------------------- |
| Scope   | Global + Local      | Only Local           |
| Style   | Full declaration    | Shortcut             |
| Type    | Explicit / Inferred | Always Inferred      |
| Usage   | Everywhere          | Inside function only |


var ব্যবহার করলে variable declare করে পরে value দেয়া যায়, আর := ব্যবহার করলে এক লাইনে declare + assign করতে হয় এবং **এটি শুধু function এর ভিতরে কাজ করে।**

### Variable Naming Rules

- A variable name must start with a letter or an underscore character (\_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and \_ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

For multi-word variable name:

1. Camel Case
   - myVaribaleName = value
2. Pascal Case
   - MyVariableName = value
3. Snake Case
   - my_variable_name = value
