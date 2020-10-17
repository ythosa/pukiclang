# pukiclang
![pukiclang](https://github.com/Ythosa/pukiclang/workflows/pukiclang/badge.svg)

## Description
PukicLang is interpreter for PukicLang programming language

## Syntax

### String, Integer, Bool variables: 
```
let age = 228;
let name = "Pukic :)";
let result = 10 * (20 / 2);
```

### Arrays:
```
let myArray = [1, 2, 3, 4, 5];
```

### Hashmaps:
```
let zalupa = {"name": "Thorsten", "age": 28};
```

### Functions:
#### Simple function:
```
let add = fn(a, b) { return a + b; };
```

#### Function without return statement:
```
let add = fn(a, b) { a + b; };
```

#### Recursive function:
```
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
```

#### High order function:
```
let twice = fn(f, x) {
  return f(f(x));
};

let addTwo = fn(x) {
  return x + 2;
};

twice(addTwo, 2); // => 6
```
