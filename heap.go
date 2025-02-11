package main

import (
	"golang.org/x/exp/constraints"
)

type HeapInterface[T constraints.Ordered] interface {
	Peek() T      // returns value at top of heap. O(1)
	RemoveTop()   // removes element at top of heap. O(logn)
	Insert(val T) // inserts data val, increases size by 1. O(logn)
	Empty() bool  // returns whether heap is empty. O(1)
	Size() int    // returns size of heap. O(1)
	Sorted() []T  // returns the elements in the heap in sorted order. O(nlogn)
}

// MaxHeap should implement HeapInterface; write a constructor & methods to complete it
type MaxHeap[T constraints.Ordered] struct {
	arr []Pair[T]
}

// MinHeap should implement HeapInterface; write a constructor & methods to complete it
type MinHeap[T constraints.Ordered] struct {
	arr []Pair[T]
}
