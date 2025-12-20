# On keeping your code happy path left aligned

Keeping your happy path on the leftmost indentation level helps to make your
code easier to follow and reason about, b/c the mental model of what's we are
trying to accomplish (AKA the behavior) is kept as linear/procedural as possible.

What do I mean by that, take the following example (taken from 100 go mistakes,
page 8).

```go
func join(s1, s2 strin, max int) (string, error) {
    if s1 == "" {
        return "", errors.New("s1 is empty")
    } else {
        if s2 == "" {
            return "", errors.New("s2 is empty")
        } else {
            concat, err := concatenate(s1, s2)
            if err != nil {
                return "", err
            } else {
                if len(concat) > max {
                    return concat[:max], nil
                } else {
                    return concat, nil
                }
            }
        }
    }
}
```

As code becomes more horizontal(indented), it forces you to keep more and more
context in your mind, the further you advance to the right the more branching
paths you need to keep track of.

On the other hand let's look at that same code refactored to try and keep
it as left aligned as possible.

```go
func join(s1, s2 strin, max int) (string, error) {
    if s1 == "" {
        return "", errors.New("s1 is empty")
    }

    if s2 == "" {
        return "", errors.New("s2 is empty")
    }

    concat, err := concatenate(s1, s2)

    if err != nil {
        return "", err
    }

    if len(concat) > max {
        return concat[:max], nil
    }

    return concat, nil
}
```

In this cleaner/better version of the function it is pretty easy to reason about
what's happening and any given moment of the func behavior, and discard code
from our mental space, if you follow it into a branch you can discard the rest
of the code. And if you keep going down you can discard all code indented
deeper that the path you are following.

In other words we got rid of the cognitive load of keeping track of multiple
branching paths at any given moment in the function lifecycle.

`Mat Rier` suggests int `The Go Time` podcast ->

| Align the happy path to the left; you should be able to scan down one column
| to see the expected execution flow.

All our error paths fall out (nested deeper) from our main execution flow.

Good rules of thumb and things to keep in mind ->

1. Align happy path to the left most indentation level (helps identify possible
   refactors into their own functions if this is not possible/easy to do)
2. When a `if/else` block returns, omit indenting logic further, refactor into a
   guard statement, this makes your code easier to read and reason about, keeps
   code left align, it's cleaner and easier to discern what the main purpose of
   your function is.
3. As mentioned in the first point if you find yourself in a position where this
   is not possible is a good indication/smell that you might be trying to do too
   much in a single function and should consider extracting some of that logic
   into its own functions to keep that one leaner.

As with everything in code each situation is different and you should reason
about it on a case per case basis, but the above are a really good guideline to
follow to help us write better Go code.
