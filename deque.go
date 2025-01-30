package main

// Input Restricted Queues
// In Input Restricted Queues, insertion takes place only from the rear end, deletion can take place from both ends.
// Output Restricted Queues
// In Output Restricted Queues, the deletion takes place from only the front, however, insertion can take place from both ends.
type DequeInterface[T any] interface {
	Front() T        // returns first item in deque. O(1)
	Back() T         // returns last item in deque. O(1)
	PushFront(val T) // adds val to front of queue, increases size by 1. O(1)
	PushBack(val T)  // adds val to back of queue, increases size by 1. O(1)
	PopFront()       // removes item from front of queue, decreases size by 1. O(1)
	PopBack()        // removes item from back of queue, decreases size by 1. O(1)
	Empty() bool     // returns whether queue is empty. O(1)
	Size() int       // returns number of elements in queue. O(1) or O(n) depending on implementation
}

type Deque[T any] struct {
	list DoublyLinkedListInterface[T]
}

func (d *Deque[T]) Front() T {
	return d.list.Head().Data
}

func (d *Deque[T]) Back() T {
	return d.list.Tail().Data
}

func (d *Deque[T]) PushFront(val T) {
	d.list.InsertAtFront(val)
}

func (d *Deque[T]) PushBack(val T) {
	if d.list.Empty() {
		return
	}
	d.list.InsertAtEnd(val)
}

func (d *Deque[T]) PopFront() {
	if d.list.Empty() {
		return
	}
	d.list.RemoveAtFront()
}

func (d *Deque[T]) PopBack() {
	d.list.RemoveAtEnd()
}

func (d *Deque[T]) Empty() bool {
	return d.list.Empty()
}

func (d *Deque[T]) Size() int {
	return d.list.Size()
}
