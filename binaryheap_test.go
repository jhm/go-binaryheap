package binaryheap

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func cmp(a, b interface{}) bool {
	return a.(int) > b.(int)
}

func elements() []interface{} {
	rand.Seed(time.Now().UnixNano())
	result := make([]interface{}, rand.Intn(19)+1)
	for i := range result {
		result[i] = rand.Intn(100)
	}
	return result
}

func validate(t *testing.T, h *BinaryHeap, i int) {
	t.Helper()
	l := left(i)
	validateChildren(t, h, i, l)
	validateChildren(t, h, i, l+1)
}

func validateChildren(t *testing.T, h *BinaryHeap, p int, c int) {
	t.Helper()
	if c < h.Len() {
		if h.cmp(h.heap[c], h.heap[p]) {
			t.Errorf("heap invariant invalidated\n[%2d]: %d\n[%2d]: %d\n", p, c, h.heap[p], h.heap[c])
			return
		}
		validate(t, h, c)
	}
}

func TestIsEmpty(t *testing.T) {
	h := New(cmp)
	if !h.IsEmpty() {
		t.Fatalf("IsEmpty returned false on an empty heap")
	}
	h.Push(1)
	if h.IsEmpty() {
		t.Fatalf("IsEmpty returned true on a non-empty heap")
	}
}

func TestEmptyPeek(t *testing.T) {
	h := New(cmp)
	if got := h.Peek(); got != nil {
		t.Errorf("Peek returned %v expected nil", got)
	}
}

func TestEmptyPop(t *testing.T) {
	h := New(cmp)
	if got := h.Pop(); got != nil {
		t.Errorf("Pop returned %v expected nil", got)
	}
}

func TestLen(t *testing.T) {
	xs := elements()
	h := NewWithElements(cmp, xs...)
	want := len(xs)
	if got := h.Len(); got != want {
		t.Errorf("Len returned wrong output\n got: %d\nwant: %d", got, want)
	}
}

func TestPeekAndPop(t *testing.T) {
	xs := elements()
	h := NewWithElements(cmp, xs...)
	sort.Slice(xs, func(i, j int) bool { return xs[i].(int) > xs[j].(int) })

	for _, want := range xs {
		if got := h.Peek().(int); got != want {
			t.Errorf("Peek returned wrong output\n got: %d\nwant: %d", got, want)
		}
		if got := h.Pop().(int); got != want {
			t.Fatalf("Pop returned wrong output\n got: %d\nwant: %d", got, want)
		}
		validate(t, h, 0)
	}
}

func TestPush(t *testing.T) {
	h := New(cmp)
	for _, n := range elements() {
		h.Push(n)
		validate(t, h, 0)
	}
}

func TestPushAll(t *testing.T) {
	h := New(cmp)
	h.PushAll(elements())
	validate(t, h, 0)
}
