package priorityqueue

import (
	"golang.org/x/exp/constraints"
)

type PriorityQueueInterface[T constraints.Ordered, V any] interface {
	Front() Pair[T]            // returns first item in queue. O(1)
	Enqueue(priority T, val V) // adds val to queue with priority, increases size by 1. O(1)
	Dequeue()                  // removes the highest priority item form queue, decreases size by 1. O(1)
	Empty() bool               // returns whether queue is empty. O(1)
	Size() int                 // returns number of elements in queue. O(1)
}

// PriorityQueue should implement PriorityQueueInterface; write a constructor & methods to complete it
type PriorityQueue[T constraints.Ordered, V any] struct {
	heap MaxHeap[T]
}
