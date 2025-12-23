# On Type Embedding

In Go, type embedding allows a struct to include the fields
of another type as if they were its own fields.

Hence we use embedding to promote the fields and methods of the embedded
type to be accessible directly from the parent struct, e.g.

```go
type Person struct {
    Name string
    Email string
}

type Employee struct {
    EmployeeID int
    // We embed the Person struct by including it without declaring a name
    Person
}

// We can access the fields of the embedded struct as if they were part of the
// parent struct without referencing the Person struct

Employee.Name // -> This is equivalent Employee.Person.Name
Employee.Email // -> This is equivalent Employee.Person.Email
```

Embedding is a powerful tool that can be used to promote composability, reduce
boilerplate and reuse code, without having to navigate through a sea of
forwarding fields/methods.

On the other hand, it can also lead to unintentional behavior and bugs, so let's
take a look at some of the pitfalls of embedding.

## Foot 🦶 Gun 󰜃 #1 Unintentional Public Access

One of the most common mistakes with embedding is given public access to
implementation internals that should remain private.

A common example is adding a mutex to a struct, e.g.

```go
type SharedCtx struct {
    async.Mutex
    data []byte
}

func (s *SharedCtx) Read(p []byte) (n int, err error) {
    s.Lock()
    defer s.Unlock()
    n := copy(p, s.data)
    return n, nil
}

// Although this allows to lock/unlock the mutex in internal methods,
// it also exposes the Mutex to the public interface, exposing internals.
// Users of our struct type could do something like this:

shared := SharedCtx{}
shared.Lock()

// This can lead to unintentional behavior and bugs.
// This is a example of an embedding bad practice.
```

> **Bad Practice**: Unintentionally exposing private implemntation internals
> to the public interface.

This can become a problem since users might assume this is part public API,
and not just an unintentional side effect of the internal implementation.

**How can we avoid this?** We can use a `private` field to hide the internals
instead of just embedding the type, that's how we fix the issue.

```go
type SharedCtx struct {
    mu sync.Mutex
    data []byte
}
```

## Foot 🦶 Gun 󰜃 #2 Thinking of embedding as substitute for inheritance

> Embedding is about composition, not inheritance or subclassing.

**Composition is a design pattern that allows you to build complex objects
from simpler ones.**

Embedding is about composability of data structures and types into more complex
ones, it doesn't inherit fields or methods from the embedded type, nor does the
type doing the embedding become a subclass.

The embedded types are still there as receivers of the promoted fields/methods,
and are also just regular fields it the type doing the embedding. It gives us
syntactic sugar to easier access and reduces boilerplate by not having navigate
additional forwarding fields/methods.

## Foot 🦶 Gun 󰜃 #3 Dependency of the current contract and future updates

When embedding, we need to be mindful of future changes in the embedded types,
since changes might be out of our control and break our code.

An example would be adding additional fields/methods that we don't want to
expose, but we unintentionally do through embedding.

Another one is accidentally shadowing an embedded type method (or field) that
is being used by the parent struct, by adding one of our own to the parent.
