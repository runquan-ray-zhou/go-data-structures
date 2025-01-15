package main

type SingleLinkNode[T any] struct {
	Data T
	Next *SingleLinkNode[T]
}

type SinglyLinkedListInterface[T any] interface {
	InsertAfter(val T, prev *SingleLinkNode[T]) // create new node with data val after node prev, increase size by 1. O(1)
	RemoveAfter(prev *SingleLinkNode[T])        // remove node after node prev, decrease size by 1. O(1)
	InsertAtFront(val T)                        // create node with data val at front of list, increase size by 1. O(1)
	RemoveAtFront()                             // remove node at front of list, decrease size by 1. O(1)
	InsertAtEnd(val T)                          // create node with data val at end of list, increase size by 1. O(n)
	RemoveAtEnd()                               // remove node at end of list, decrease size by 1. O(n)
	Head() *SingleLinkNode[T]                   // return first node in list. O(1)
	Tail() *SingleLinkNode[T]                   // return last node in list. O(n)
	Empty() bool                                // returns whether list is empty. O(1)
	Size() int                                  // returns number of elements in list. O(1) or O(n) depending on implementation
}

type SinglyLinkedList[T any] struct {
	head *SingleLinkNode[T]
	tail *SingleLinkNode[T]
	// you can choose whether to store the size by uncommenting the following line
	size int
}

func (s *SinglyLinkedList[T]) InsertAfter(val T, prev *SingleLinkNode[T]) {
	newNode := &SingleLinkNode[T]{Data: val}
	if prev.Next == nil {
		s.tail = newNode
		prev.Next = newNode
		s.size++
	} else {
		restOfList := prev.Next
		newNode.Next = restOfList
		prev.Next = newNode
		s.size++
	}

}

func (s *SinglyLinkedList[T]) RemoveAfter(prev *SingleLinkNode[T]) {
	if s.head == nil || prev == nil || prev == s.tail {
		return
	}
	prev.Next = prev.Next.Next
	s.size--
}

func (s *SinglyLinkedList[T]) InsertAtFront(val T) {
	newNode := &SingleLinkNode[T]{Data: val}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.Next = s.head.Next
		s.head = newNode
	}
	s.size++
}

func (s *SinglyLinkedList[T]) RemoveAtFront() {
	if s.head == nil {
		return
	}
	s.head = s.head.Next
	s.size--
}

func (s *SinglyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &SingleLinkNode[T]{Data: val}
	currNode := s.head
	if currNode.Next == nil {
		currNode.Next = newNode
	}
	for currNode != nil && currNode.Next != nil {
		if currNode.Next.Next == nil {
			currNode.Next = newNode
		}
		currNode = currNode.Next
	}
	s.tail = newNode
	s.size++
}

func (s *SinglyLinkedList[T]) RemoveAtEnd() {
	if s.head == s.tail {
		s.head = nil
	}
	currNode := s.head
	for currNode != nil && currNode.Next != nil {
		if currNode.Next.Next == nil {
			currNode.Next = currNode.Next.Next
		}
		currNode = currNode.Next
	}
	s.tail = currNode
	s.size--
}

func (s *SinglyLinkedList[T]) Head() SingleLinkNode[T] {
	return *s.head
}

func (s *SinglyLinkedList[T]) Tail() SingleLinkNode[T] {
	return *s.tail
}

func (s *SinglyLinkedList[T]) Empty() bool {
	return s.head == nil
}

func (s *SinglyLinkedList[T]) Size() int {
	return s.size // size type is not a pointer
}
