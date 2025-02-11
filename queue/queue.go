package queue

import "fmt"

// FIFO
// are we using fixed size circular arrays or dynamic size circular arrays?
type QueueInterface[T any] interface {
	Front() T      // returns first item in queue. O(1)
	Enqueue(val T) // adds val to end of queue, increases size by 1. O(1)
	Dequeue()      // removes item from front of queue, decreases size by 1. O(1)
	Empty() bool   // returns whether queue is empty. O(1)
	Size() int     // returns number of elements in queue. O(1) or O(n) depending on implementation
	StringifyQueue() string
}

type Queue[T any] struct { //Head of singly linked list is front of queue and tail is back of queue
	list SinglyLinkedListInterface[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: NewSinglyLinkedList[T]()}
}

// first is the first element enqueued into the array or when the array was instantiated.
// Alternate implementation of a queue using a circular array
type AlternateQueue[T any] struct {
	arr   []T
	first int
	last  int
	size  int
}

func NewArrayQueue[T any](capacity int) *AlternateQueue[T] {
	return &AlternateQueue[T]{
		arr:   make([]T, capacity), // have capacity, meaning it is for fixed size
		first: 0,
		last:  0,
		size:  0,
	}
}

func (q *Queue[T]) Front() T {
	if q.list.Empty() {
		fmt.Println("Queue is empty")
	}
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
	if aq.Empty() {
		fmt.Println("Queue is empty")
	}
	return aq.arr[aq.first]
}

func (aq *AlternateQueue[T]) Enqueue(val T) {
	if aq.size == len(aq.arr) {
		fmt.Println("Queue is full")
		return
	}

	aq.last = (aq.last + 1) % len(aq.arr)
	aq.arr[aq.last] = val
	aq.size++
}

func (aq *AlternateQueue[T]) Dequeue() {
	if aq.Empty() {
		fmt.Println("Queue is empty")
		return
	}

	aq.first = (aq.first + 1) % len(aq.arr)
	aq.size--
}

func (aq *AlternateQueue[T]) Empty() bool {
	return aq.size == 0
}

func (aq *AlternateQueue[T]) Size() int {
	return aq.size
}

func (aq *AlternateQueue[T]) StringifyQueue() string {
	if aq.Empty() {
		return "Queue is empty"
	}

	result := ""
	count := 0

	for i := aq.first; count < aq.size; i = (i + 1) % len(aq.arr) {
		result += fmt.Sprintf("%v", aq.arr[i])

		if count < aq.size-1 {
			result += " -> "
		}
		count++
	}
	return result
}
