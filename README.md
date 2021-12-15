# go-binaryheap

A slice backed binary heap where the order can be customized by a comparison
function. The main branch now requires go 1.18 because the heap makes use of
generic type parameters. For a version that works on Go 1.17 and below see the
v1.0 tag.

## Usage

```golang
// Construct a new max heap containing ints.
h := binaryheap.New(func(a, b int) bool { return a > b })

// Add an int to the heap.
h.Push(1)

// Add multiple ints to the heap.
h.PushAll(2, 3)

// Retrieve the top item.
a, found := h.Peek()

// Retrieve the top item and remove it from the heap.
b, found := h.Pop()
```

## Future Changes

The tests will most likely be updated to use the new fuzzing API released in Go 1.18.
