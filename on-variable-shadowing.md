# On variable shadowing

Unintended variable shadowing can lead to unexpected behavior in Go.

For that reason it's important to limit the scope of variables as much as possible.

Take a look at the following example:

```go
func main() {
 var result int

 if shouldFetch1() {
  result, err := FetchCount1()
  if err != nil {
   log.Fatal("failed to fetch count 1")
  }
 } else {
  result, err := FetchCount2()
  if err != nil {
   log.Fatal("failed to fetch count 2")
  }
 }

 fmt.Println("The result is ->", result)
}
```

Here the result in the inner `if` block is unintentionally being created and
scoped to the `if` block, instead of being assigned to the already existing one,
causing Unintended behavior, the variable in the `main func` scope will never be updated.

How can we make sure the intended behavior happens then? We create an inner
placeholder variable, and then we assign it to the main variable.

```go
var result int
if shouldFetch1 {
    r, err := FetchCount1()
    if err != nil {
        log.Fatal("failed")
    }
    result = r
}
```

We can also just assign inside the `if` block and declare the variables above that
scope.

```go
var result int
var err error
if shouldFetch1 {
    result, err = FetchCount1()
    if err != nil {
        log.Fatal("failed")
    }
} else {
    // Same logic
}
```

This second option also allows to write less code, b/c we can extract the error
handling outside the branching `if/else` code block.

```go
var result int
var err error
if shouldFetch1 {
    result, err = FetchCount1()
} else {
    result, err = FetchCount2()
}

if err != nil {
    log.Fatal("failed")
}
```

This is fine int this very small example, but I whenever possible I prefer
to stick to creating scoped variables (first example), keeping the scope as
small as possible, and handling errors as close as possible to where they
happen, but as with everything in code context is king, and you should decide
which is the best option for your particular case.
