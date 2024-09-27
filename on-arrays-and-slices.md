# On Arrays and Slices

A brief description of the behavior of Array and Slices in the Go Programming language.

## Lessons Learned

Arrays and Slices behave differently when creating new variables from
existing ones. A brief description of the differences that have given me a
gotcha moment in some occasion.

### Discrete(Array) vs Composite(Slice) types

Arrays behave as discrete types, creating a new var in Go allocates a certain
amount of memory, so when we create an Array var from another variable go allocates
enough memory to copy over the whole chunk from the original array to the new location.

This is easiar to visualize if we break apart the operations instead of doing it in a
single expresion:

![image](https://github.com/user-attachments/assets/c305670b-a79b-4ad9-bfb7-5cd449925ae5)

Then we do the assignment:

![image](https://github.com/user-attachments/assets/38a20f11-9698-4e78-ab7a-7c29ae564a21)

This is the same than doing it all at once like so:

![image](https://github.com/user-attachments/assets/cf08c2f6-3027-4fc2-b59b-59e572782eae)

This might be confusing when thinking and comparing this behavior with the behavior of slices,
which are a composite type, that means they are composed of 3 fields, a len, a cap and a pointer
to the memory location of the backing array for the slice elements.

![image](https://github.com/user-attachments/assets/248cec89-6ba9-4600-8dc5-2ca8f4176864)

What that means is that creating a new variable from another slice var,
will copy the composite structure (struct) of the slice, along with the pointer,
they will share the same backing array (pointer) and memory location, 
in memory it would look something like this.

![image](https://github.com/user-attachments/assets/bfbdac7a-8433-4508-8fc4-ca62f52abb27)

This is the reason why changing the elements in one will affect the other,
and also what happens when you do sub-slices, which have the exact same behavior,
and changing the subslice will change the shared backing array.

This changes if you start appending to the slice, and at some point when there's not
enough capacity in the backing array, a new one will be created that has enough capacity
to hold the new elements. The behavior of `append()` is to allocate a new array in memory,
then update the pointer and set it to this new location in memory. It might look something
like this:

![image](https://github.com/user-attachments/assets/b8de3bcb-a934-43bb-ad03-d474bf9e3b13)

Once that happens the two slices no longer share the backing array in memory and updating
elements in either won't affect the other, since the original one will still point to the
first backing array.

## To summarize:

### Array Behavior when creating a new var from an existing array var

1. Allocates new memory and copy the whole chunk of memory from the original
   array (where the elements are stored).
2. Modifying elements in the new var updates a diff memory location, the two
   vars do not share the same memory location.

### Slice Behavior when creating a new var from an existing slice var

1. Allocates enough mem for the capacity, len and a pointer.
2. Shares the same mem location for the backing array.
3. Once we start appending to either slice, and not enough capacity, a new bigger
   backing backing array with more capacity is created, and the pointer in the slice
   updated to point to this new memory location.

