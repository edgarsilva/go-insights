# On Interfaces

Interfaces are one of the most useful and powerful tools in the Go programming
language arsenal, but as with anything else, too much of a good thing can become
bad and cause more issues than solving.

The real problem is overusing them, filling our code with **unnecessary abstractions**
where simply passing a concrete type would've been the right approach.

But let's backtrack for a little bit, `What's an Interface?` Well they are
mostly the same than in any other programming language:

> `An intarface is way to define behavior of an object/class/struct`

What makes Go Interfaces different than in most other languages is
that `they are satisfied implicitly`.

> `Interfaces in Go are satisfied implicitly`

What this means is that unlike other languages (like Java), where you usually
"Implement" an interface, in Go any object that satisfies the requirements is
also identified as implementing the interface.

Take the `io.Reader` interface for example, which contains a single
method definition ->

```go
type Reader Interface {
    Read(p []byte) (n int, err error)
}
```

Now let's take a look at the following function that accepts an `io.Reader`.

```go

type StrReader struct {
    str string
}
func (s StrReader) Read(buf []byte) (int, error) {
    // Implement Read method
}

type BufferReader struct {
    buf []byte
}
func (s BufferReader) Read(buf []byte) (int, error) {
    // Implement Read method
}

type SerialReader struct {
    data myStruct
}
func (s SerialReader) Read(buf []byte) (int, error) {
    // Implement Read method
}

// This function will accept any of the above readers
// without our implementation ever having to know the
// concrete type, and without the consumers ever having
// to explicitly implement the interface.
func printSrcToStdout(src io.Reader) {
    // Implementation
}
```

This way we achieve several things, decouple producer/client from
having to know inner library implementation, API contract can be
fulfilled without forcing concrete types, useful to avoid circular
dependencies, promotes dependency injection by design, makes testing
easier since interfaces can be easily mocked/stubbed.

A test example of this could be as follows:

**Note** Example taken from `100 Go Mistakes`

```go
func TestCopySrcToDst(t *testing.T) {
    cost input = "foo"
    src := strings.NewReader(input)
    dst := bytes.NewBuffer(make([]byte, 0))

    err := CopySrcToDst(dst, src)
    if err != nil {
        t.FailNow()
    }

    got := dst.String()
    if got != input {
        t.Errorf("Expected: %s, got: %s", input, got)
    }
}
```

You can see here that the behavior/implementation of the `CopySrcToDst` func
does not care what the object passed is, provided it satisfies the interface
requirements it will be accepted, it could be a file, a stream, a network or
hardware interface, etc.

When creating interfaces the smaller the better, the `io.Reader` above and
`io.writer` are good examples of this, both interfaces define a single method.

TL;DR when defining interfaces the smaller the better.

Rob Pike one of the creators of go, is famously quoted saying ->

> **The bigger the interface the weaker the abstraction.**
>
> - By Rob Pike

Adding methods to an interface can drastically impact how reusable it is, an
interface must be as simple as possible, if in need of more complex behavior
we can always use composability to define it.

> We can re-use interfaces to compose more complex behaviors e.g. ReadWriter

```go
type ReadWriter interface {
    Reader
    Writer
}
```

## When and Why to use Interfaces?

There are at least four good reasons to use interfaces in Go:

1. Common behavior
2. Decoupling
3. Dependency Injection
4. Restricting behavior in our methods/functions

There are also some other good reasons to use interfaces, but the above are the
most common ones. Other reasons.

- Pluggable behavior
- Testing
- External systems/adapters

Let's take a look at a good example of the first one, a `Sortable` interface.

```go
type Sortable interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

func IsSorted(s Sortable) bool {
    n := s.Len()
    for i := 0; i < s.Len(); i++ {
        if i < n-1 && s.Less(i, i+1) {
            return false
        }
    }
    return true
}
```

This is a good example of a common behavior that we can use across our codebase
that has high potential for reuse.

## Interfaces No Bueno Practices (AKA Bad Practices)

### Interface Pollution Foot 🦶 Guns 🔫

Most common is overusing them and creating unnecessary abstractions. You should
always start with the simplest solution (concrete types like structs) and
refactor it into a more complex one when the need arises.

> Start with concrete types, refactor into interfaces when needed.

In Go it is almost always a mistake to start with an abstraction, in this cause
an `interface`, forgetting one of the main tenants of programming an
abstraction:

> Abstractions should be discovered, not created.

What the means is that we shouldn't start creating abstractions in our code if
there is no immediate and clear need for it.

The `100 Go Mistakes` puts it very clearly in a single sentence ->

> We should create an interface when we need it, not when we foresee that we
> could need it.
>
> by Teiva Harsanyi (100 Go Mistakes book)

The main problem if we overuse interfaces is that we make the code very hard to
follow, the flow of the code gets way more complex adding useless levels of
indirection that don't add any value to the code.

> Don't design with interfaces, discover them.
>
> by Rob Pike

### Where should interfaces live?

> Interfaces should live on the **Consumer** side.

But where is this?

This is where the interface is being used, not where the concrete type that
meets the interface requirements is defined, we call that the **Producer** side.

> Don't put interfaces in the **Producer** side.

An exaple of this, in one application we have a Service Auth `auth/service.go`,
this service accepts `Strategies` that satisfy the `auth.Strategy`

```go
// auth/service.go
type Strategy interface {
    Authenticate(c *fiber.Ctx) (model.User, error)
}
```

Then the concrete implementation of each strategy that leaves in it's own
package must define an `Authenticate` method, e.g.

```go
// auth/local/strategy.go
type LocalStrategy struct {
    Ctx     *fiber.Ctx
    service LocalStrategyService
}
func (ls LocalStrategy) Authenticate(c *fiber.Ctx) (model.User, error) {
    // Implementation
}
```

We could later decide we want more Authentication strategies, now we can easily
add a new one, without having to change the `auth/service.go` file, it could
look something like this:

```go
// auth/logto/strategy.go
type LogtoStrategy struct {
 Ctx      *fiber.Ctx
 Logto    *logto.LogtoClient
 service  LogtoStrategyService
}
func (s *LogtoStrategy) Authenticate(c *fiber.Ctx) (model.User, error) {
    // Implementation
}
```

> TL;DR You should put interfaces in the consumer side (where the interface is
> being used), not the producer side (where the concrete type is defined).

### Returning Interfaces from functions/methods

Short answer `don't do it, it's a bad idea 99.999%`.

In most cases you should never return an interface from a function/method, it's
considered a bad practice in Go, and it's a sign of bad design. Returning an
interface restricts flexibility and forces clients to a single abstraction.

> Be conservative in what you do, be liberal in what you
> accept from others.
>
> Postel's Law

In go this means return structs instead of interfaces, and accept Interfaces if
possible and the use case warrants it.

There's a saying in Go that goes like this ->

> Accept interfaces and return concrete types.
>
> Efficient Go (probably)

Of course there are exceptions to this rule, but in general it's a good rule
of thumb. The most common one is the `Error Type` which is an interface often
returned by many of your functions.
