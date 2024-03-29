package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is a type that implements heap.Interface and holds integers.
type IntHeap []int

// Len returns the number of elements in the IntHeap.
func (h IntHeap) Len() int { return len(h) }

// Less returns true if element at i is less than element at j.
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap swaps the elements at i and j.
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *IntHeap) Push(x any) {
	// x is asserted to be of type int, since IntHeap is a slice of ints.
	*h = append(*h, x.(int))
}

// Pop removes and returns the element at the end of the heap.
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // Get the last element of the slice.
	*h = old[0 : n-1] // Remove the last element from the slice.
	return x
}

func main() {
	h := &IntHeap{2, 1, 5, 7, 4, 8, 9, 6}
	heap.Init(h) // This will arrange the IntHeap to satisfy the heap invariant.

	heap.Push(h, 3)
	fmt.Printf("Min-heap after Push(3): %v\n", *h) // [1 2 5 3 4 8 9 7 6]

	min := heap.Pop(h)
	fmt.Printf("Min-heap after Pop(): %v, Popped element: %v\n", *h, min) // [2 3 5 6 4 8 9 7],  1
}
