# On Interfaces

Interfaces are one of the most useful and powerful tools in the Go programming
language arsenal, but as with anything else, too much of a good thing can become
bad and cause more issues than solving.

The real problem is overusing them, filling our code with **unnecessary abstractions**
where simply passing a concrete type to a function would've been the right
approach.

But `What's an Interface?` They are same than in any other programming language
(maybe just not Typescript), `A way to define behavior of an
object/class/struct`, depending on the lang being used.

What makes Go Interfaces different than in most other languages is
that `they are met/honored implicitly`.

### Interfaces in Go are met implicitly

What this means is that unlike other languages (like Java), where you usually
"Implement" an interface, in go any object that meets the requirements is
also identified as implementing the interface.

Let's see a simble example using `io.Reader` interface which contains a single
method ->

```go
type Reader Interface {
    Read(p []byte) (n int, err error)
}
```

We could write the following function (a naive implemetation for
demostration purposes only) to test this.

```go
package main

import (
 "fmt"
 "io"
 "log"
)

func main() {
 str := "Hello World of Golang!"
 src := StrReader{str: &str}
 dst := &StrWriter{}

 printSrcToDst(src, dst)

 fmt.Println("Full written dst:", string(dst.data))
}

type StrReader struct {
 str *string
}

func (s StrReader) Read(buf []byte) (int, error) {
 if len(*s.str) == 0 {
  return 0, io.EOF
 }

 n := copy(buf, *s.str)
 *s.str = (*s.str)[n:]

 return n, nil
}

type StrWriter struct {
 data []byte
}

func (w *StrWriter) Write(buf []byte) (int, error) {
 fmt.Println("writing ->", string(buf))
 w.data = append(w.data, buf...)
 return len(buf), nil
}

func printSrcToDst(src io.Reader, dst io.Writer) {
 buf := make([]byte, 8)

 for {
  n, err := src.Read(buf)
  if n > 0 {
   _, err := dst.Write(buf[:n])
   if err != nil {
    log.Fatal(err)
   }
  }
  if err == io.EOF {
   break
  }
  if err != nil {
   panic(err)
  }
 }
}
```

Here we could replace `StrReader` or `StrWriter` with any other object that
implements the `Read` or `Write` functions like `Response` and `Request`,
or a file, or other network implementations.

This is really powerful b/c it let's you create functions that expect an
argument to meet an interface, and the consumers don't need to "implement" the
interface explicitly, as long as the passed object meets the requirements it
will be accepted, this decouples client code from your implementation and has
other uses like avoiding circular references and simplifies and promotes
dependency injection by design.

It also makes testing easier by the very nature of receiving interfaces instead
of concrete types, like the aformentioned `io.Reader/Writer` vs using real files.

We could use the Go standard library helper methods to create Readers/Writers
like so ->

**Note** Examble taken from `100 Go Mistakes`

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

You can see here that the behavior/implementation of the method does not care
what the object is, that go identifies it as implementing the interface without
us ever having to directly implement it.

When creating interfaces the smaller the better, Rob Pike one of the creators of
go, famously said this ->

## **The bigger the interface the weaker the abstraction. By Rob Pike

Adding methods to an interface can drastically reduce how re-usable it is, an
interface must be as simple as possible, we can always use composability os
interfaces to define more complex behaviors.

### We can re-use interfaces to compose more complex behaviors e.g. ReadWriter

```go
type ReadWriter interface {
    Reader
    Writer
}
```
