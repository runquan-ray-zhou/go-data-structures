package main

// FIFO
type QueueInterface[T any] interface {
	Front() T      // returns first item in queue. O(1)
	Enqueue(val T) // adds val to end of queue, increases size by 1. O(1)
	Dequeue()      // removes item from front of queue, decreases size by 1. O(1)
	Empty() bool   // returns whether queue is empty. O(1)
	Size() int     // returns number of elements in queue. O(1) or O(n) depending on implementation
}

type Queue[T any] struct { //Head of doubly linked list is front of queue and tail is back of queue
	list SinglyLinkedListInterface[T]
}

// Alternate implementation of a queue using a circular array
type AlternateQueue[T any] struct {
	arr   []T
	first int
	last  int
}

func (q *Queue[T]) Front() T {
	return q.list.Head().Data
}

func (q *Queue[T]) Enqueue(val T) {
	q.list.InsertAtEnd(val)
}

func (q *Queue[T]) Dequeue() {
	q.list.RemoveAtFront()
}

func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}

func (q *Queue[T]) Size() int {
	return q.list.Size()
}

func (aq *AlternateQueue[T]) Front() T {
	return aq.arr[aq.first]
}

func (aq *AlternateQueue[T]) Enqueue(val T) {
}

func (aq *AlternateQueue[T]) Dequeue() {
}

func (aq *AlternateQueue[T]) Empty() bool {
}

func (aq *AlternateQueue[T]) Size() int {
}
