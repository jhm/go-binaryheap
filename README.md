# go-binaryheap [![Actions Status](https://github.com/jhm/go-binaryheap/workflows/Main/badge.svg)](https://github.com/jhm/go-binaryheap/actions)

A slice backed binary heap where the order can be customized by a comparison
function.

## Usage

```golang
// Construct a new max heap containing ints.
h := binaryheap.New(func(a, b interface{}) bool { return a.(int) > b.(int) })

// Add an int to the heap.
h.Push(1)

// Add multiple ints to the heap.
h.PushAll(2, 3)

// Retrieve the top item.
a := h.Peek()

// Retrieve the top item and remove it from the heap.
b := h.Pop()
```

## Future Changes

Upon the release of Go 1.18 the heap, and its API, will make use of generic type
parameters and the tests will most likely be updated to use the new fuzzing API.
