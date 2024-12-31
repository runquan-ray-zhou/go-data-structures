package godatastructures

type QueueInterface[T any] interface {
	Front() T      // returns first item in queue. O(1)
	Enqueue(val T) // adds val to end of queue, increases size by 1. O(1)
	Dequeue()      // removes item from front of queue, decreases size by 1. O(1)
	Empty() bool   // returns whether queue is empty. O(1)
	Size() int     // returns number of elements in queue. O(1) or O(n) depending on implementation
}

type Queue[T any] struct {
	list SinglyLinkedListInterface[T]
}

// Alternate implementation of a queue using a circular array
type AlternateQueue[T any] struct {
	arr []T
	first int
	last int
}
