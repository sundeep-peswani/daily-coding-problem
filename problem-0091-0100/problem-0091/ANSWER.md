The code snippet prints out `9` 10 times. This is because `i` is set to `9` after the first `for` loop ends. When every function in `functions` is executed, they all reference and return the same `i`, which is set to `9`.

To solve this problem, we can capture the `i` when creating the lambda, by changing the way we create and append the lambda to `functions`:

```
functions = []
for i in range(10):
    functions.append(lambda i=i: i)

for f in functions:
    print(f())
```

`i=i` captures `i` in the `for` loop, sets it to *another* variable `i`, which the lambda can then return.