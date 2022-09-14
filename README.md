# monkey-ball

The interpreter from Thorsten Ball's [Writing An Interpreter In Go](https://interpreterbook.com/) 
and the compiler and VM from [Writing a Compiler In Go](https://compilerbook.com/)

Samples from the unit tests:
```
### Expressions ###

>> (5 + 10 * 2 + 15 / 3) * 2 + -10
50

>> if (1 < 2) { 10 } else { 20 }
10

### Statements ###

>> let a = 5;
>> let b = a;
>> let c = a + b + 5;
>> c;
15

### Functions and Closures ###

>> let add = fn(x, y) { x + y; };
>> add(5, 5);
10

>> add(2 + 3, add(1, 2));
8

>> let newAdder = fn(x) { fn(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(7);
9

### Arrays and Builtins ###

>> len("hello world")
11

>> let a = [1, 2, 3];
>> a[0] + a[1];
3

>> a[1 + 1];
3

>> first(a)
1

>> last(a)
3

>> rest(a)
[2, 3]

>> push(a, 4)
[1, 2, 3, 4]

### Hashes ###

>> let h = {"foo": 4, "bar": 3 + 2, "foo" + "bar": len("foobar")};
>> h["foo"]
4

>> h["bar]
5

>> h["foobar"]
6
```
