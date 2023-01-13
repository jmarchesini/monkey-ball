# monkey-ball

An implementation of the Thorsten Ball's Monkey language that comes from the following sources:
- [Writing An Interpreter In Go](https://interpreterbook.com/) 
- [Writing a Compiler In Go](https://compilerbook.com/)

I have not yet implemented 
[The Lost Chapter](https://interpreterbook.com/lost/).

### Getting Started ###

Once you clone the repo, use the following to build the `monkey` executable from the `monkey-ball` directory:

`> go build -o monkey .`

Monkey can be launched using the interpreter from Ball's first book, or the compiler and vm from his second book.
To run `monkey` with the compiler/vm engine use:

`> ./monkey -engine=vm`

This is the default and is equivalent to:

`> ./monkey`

To run with the interpreter engine use:

`> ./monkey -engine=eval`

### Benchmark ###

In Chapter 10 of the compiler book, Ball supplies a benchmark to measure the difference in runtime between the interpreter and compiler in the time it takes to calculate `fibonacci(35)`.  To build the benchmark use:

`> go build -o fibonacci ./benchmark`

As above, you can supply the engine argument (defaults to `vm`).

```
> ./fibonacci
engine=vm, result=9227465, duration=5.06227534s`
```
```
> ./fibonacci -engine=eval
engine=eval, result=9227465, duration=22.992759395s`
```

### Tests ###

A large part of the code in the book is in the form of unit tests.  You can run the entire test suite with:

`> go test ./...`

### The Monkey Language ###

Samples of the Monkey language from the unit tests and run in the REPL:

#### Expressions ####
```
>> (5 + 10 * 2 + 15 / 3) * 2 + -10
50

>> if (1 < 2) { 10 } else { 20 }
10
```

#### Statements ####
```
>> let a = 5;
>> let b = a;
>> let c = a + b + 5;
>> c;
15
```

#### Functions, Scopes, and Closures ####
```
>> let add = fn(x, y) { x + y; };
>> add(5, 5);
10

>> add(2 + 3, add(1, 2));
8

>> let newAdder = fn(x) { fn(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(7);
9

>> let globalSeed = 50;
>> let minusOne = fn() { let num = 1; globalSeed - num; };
>> let minusTwo = fn() { let num = 2; globalSeed - num; };
>> minusOne() + minusTwo();
97
```

#### Arrays and Builtins ####
```
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
```

#### Hashes ####
```
>> let h = {"foo": 4, "bar": 3 + 2, "foo" + "bar": len("foobar")};
>> h["foo"]
4

>> h["bar"]
5

>> h["foobar"]
6
```
