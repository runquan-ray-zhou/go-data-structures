package main

// LIFO
type StackInterface[T any] interface {
	Top() T      // returns top item in stack. O(1)
	Push(val T)  // adds val to top of stack, increases size by 1. O(1)
	Pop()        // removes item from top of stack, decreases size by 1. O(1)
	Empty() bool // returns whether stack is empty. O(1)
	Size() int   // returns number of elements in stack. O(1) or O(n) depending on implementation
}

type Stack[T any] struct {
	list SinglyLinkedListInterface[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: NewSinglyLinkedList[T](nil),
	}
}

func (s *Stack[T]) Top() T {
	return s.list.Head().Data
}

func (s *Stack[T]) Push(val T) {
	s.list.InsertAtFront(val)
}

func (s *Stack[T]) Pop() {
	if s.list.Empty() {
		return
	}
	s.list.RemoveAtFront()
}

func (s *Stack[T]) Empty() bool {
	return s.list.Empty()
}

func (s *Stack[T]) Size() int {
	return s.list.Size()
}

// Alternate implementation of a stack using an array
type AlternateStack[T any] struct {
	arr []T
}

func NewAlternateStack[T any](initialCapacity int) *AlternateStack[T] {
	return &AlternateStack[T]{
		arr: make([]T, 0, initialCapacity),
	}
}

func (as *AlternateStack[T]) Top() T {
	return as.arr[len(as.arr)-1]
}

func (as *AlternateStack[T]) Push(val T) {
	as.arr = append(as.arr, val)
}

func (as *AlternateStack[T]) Pop() {
	if len(as.arr) == 0 {
		return
	}
	as.arr = as.arr[:len(as.arr)-1]
}

func (as *AlternateStack[T]) Empty() bool {
	return len(as.arr) == 0
}

func (as *AlternateStack[T]) Size() int {
	return len(as.arr)
}
