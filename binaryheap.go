// Package binaryheap provides an implementation of a slice backed binary heap
// where the order can be customized by a comparator function.
package binaryheap

// Comparator returns true if, and only if, 'a' has a higher priority than 'b';
// that is, 'a' should be retrieved from the heap before 'b'.
type Comparator func(a, b interface{}) bool

type BinaryHeap struct {
	heap []interface{}
	cmp  Comparator
}

// The default capacity of the slice that contains the heap elements.
const DefaultHeapCapacity = 0

// New creates a new binary heap with the given comparator and the default
// initial capacity.
func New(cmp Comparator) *BinaryHeap {
	return NewWithCapacity(cmp, DefaultHeapCapacity)
}

// NewWithCapacity creates a new binary heap with the given comparator and
// initial capacity.
func NewWithCapacity(cmp Comparator, capacity int) *BinaryHeap {
	return &BinaryHeap{
		heap: make([]interface{}, 0, capacity),
		cmp:  cmp,
	}
}

// NewWithElements creates a new binary heap with the given comparator and
// elements.
func NewWithElements(cmp Comparator, elements ...interface{}) *BinaryHeap {
	h := NewWithCapacity(cmp, len(elements))
	h.PushAll(elements...)
	return h
}

// IsEmpty returns true if there are zero elements in the heap, false otherwise.
func (h *BinaryHeap) IsEmpty() bool {
	return h.Len() == 0
}

// Len returns the number of elements in the heap.
func (h *BinaryHeap) Len() int {
	return len(h.heap)
}

// Peek returns the first element in the heap or nil if the heap is empty. The
// complexity is O(1).
func (h *BinaryHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h.heap[0]
}

// Pop returns and removes the first element in the heap. If the heap is empty
// Pop returns nil. The complexity is O(log n) where n is the number of elements
// in the heap.
func (h *BinaryHeap) Pop() interface{} {
	n := h.Len()
	if n == 0 {
		return nil
	}
	result := h.heap[0]
	h.heap[0] = h.heap[n-1]
	h.heap[n-1] = nil
	h.heap = h.heap[0 : n-1]
	h.down(0)
	return result
}

// Push adds a single element to the heap. The complexity is O(log n) where n is
// the number of elements in the heap.
func (h *BinaryHeap) Push(element interface{}) {
	h.heap = append(h.heap, element)
	h.up(h.Len() - 1)
}

// PushAll adds multiple elements to the heap. The complexity is O(n) where n is
// the number of elements in the produced heap.
func (h *BinaryHeap) PushAll(elements ...interface{}) {
	h.heap = append(h.heap, elements...)
	for i := parent(len(h.heap) - 1); i >= 0; i-- {
		h.down(i)
	}
}

// Values returns a new slice that contains all of the elements in the heap.
// The returned slice has the same order as the underlying heap slice and is NOT
// necessarily in the order that would be produced by consecutive calls to
// Pop().
func (h *BinaryHeap) Values() []interface{} {
	result := make([]interface{}, h.Len())
	copy(result, h.heap)
	return result
}

func (h *BinaryHeap) up(i int) {
	for {
		p := parent(i)
		if p == i || h.cmp(h.heap[p], h.heap[i]) {
			break
		}
		h.heap[p], h.heap[i] = h.heap[i], h.heap[p]
		i = p
	}
}

func (h *BinaryHeap) down(i int) {
	for {
		l := left(i)
		// If there is no left child stop immediately. Note that l can
		// be less than zero because of overflow.
		if l >= h.Len() || l < 0 {
			break
		}
		// Get the index of the child with the largest value.
		max := l
		if r := l + 1; r < h.Len() && h.cmp(h.heap[r], h.heap[l]) {
			max = r
		}
		// If the value at the parent index is larger than the largest
		// child value the heap invariant holds so no more swaps are
		// necessary.
		if h.cmp(h.heap[i], h.heap[max]) {
			break
		}
		h.heap[i], h.heap[max] = h.heap[max], h.heap[i]
		i = max
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}
