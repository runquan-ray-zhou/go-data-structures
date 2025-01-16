package main

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

func (s *Stack[T]) Top() T {
	topOfStack := s.list.Head()
	return topOfStack.Data
}

func (s *Stack[T]) Push(val T) {
	s.list.InsertAtFront(val)
}

func (s *Stack[T]) Pop() {
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

func (as *AlternateStack[T]) Top() T {
	topOfStack := as.arr[len(as.arr)-1]
	return topOfStack
}

func (as *AlternateStack[T]) Push(val T) {
	as.arr = append(as.arr, val)
}

func (as *AlternateStack[T]) Pop() {
	as.arr = as.arr[:len(as.arr)-1]
}

func (as *AlternateStack[T]) Empty() bool {
	return len(as.arr) == 0
}

func (as *AlternateStack[T]) Size() int {
	return len(as.arr)
}
