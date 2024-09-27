# On Arrays and Slices

A brief description of the behavior of Array and Slices in the Go Programming language.

## Lessons Learned

Arrays and Slices behave differently when creating new variables from an
existing one. A brief description of the differences that have given me a
gotcha moment in some occasion.

### Discrete(Array) vs Composite(Slice) types

They behave as discrete types, creating a new var in Go allocates a certain
amount of memory, so when we create an Array var from another one go allocates
enough memory to copy over the whole chunk of memory from the original array
to the new location.

    **-----Insert image here----**

This might be confusing with the behavior of slices, which are a composite
type, that means they are composed of 3 fields, a len, a cap and a pointer
to the memory location of the backing array for the slice elements.

    **-----Insert image here----**

What that means is that creating a new variable from another slice var,
will share the same backing array (pointer) and memory location between the
two, since remember that creating a new variable allocates enough mem for
the type ant copies the stuff over.

That means we allocate enough mem for cap, len and a pointer; one that points
to memory location of the backing array from the original slice.

Remember that this might change if you start appending to the slice, since there
behavior of `append()` is to create a new backing array in memory, that is big
enough, once there's not enough capacity to append a new element, and update
the pointer to point to this new location in memory.

Once that happens the two slices no longer share memory and updating elements in
either won't affect the other, since the original one will still point to the
original backing array pointer.

### Array Behavior when creating a new var form an existing array var

1. Allocates new memory and copy the whole chunk of memory from the original
   array (where the elements are stored).
2. Modifying elements in the new var updates a diff memory location, the two
   vars do not share the same memory location.

### Slice Behavior when creating a new var from an existing slice var

1. Allocates enough mem for cap, len and a pointer.
2. Shares the same mem location for the backing array.
3. Once we start appending to either slice, if not enough capacity, a new backing
   backing array with enough capacity is created, and the pointer updated to pointer
   to this new memory location.

### Usage/Examples

    ```go
    package main

    func main() {
        fmt.Println("pending...")
    }
    ```

## Screenshots

![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)
