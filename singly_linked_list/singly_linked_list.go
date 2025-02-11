package singlylinkedlist

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

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (s *SinglyLinkedList[T]) InsertAfter(val T, prev *SingleLinkNode[T]) {
	if prev == nil {
		return
	}
	newNode := &SingleLinkNode[T]{Data: val, Next: prev.Next}
	prev.Next = newNode
	if prev == s.tail {
		s.tail = newNode
	}
	s.size++
}

func (s *SinglyLinkedList[T]) RemoveAfter(prev *SingleLinkNode[T]) {
	if s.head == nil || prev == nil || prev.Next == nil {
		return
	}
	if prev.Next == s.tail {
		s.tail = prev
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
		newNode.Next = s.head
		s.head = newNode
	}
	s.size++
}

func (s *SinglyLinkedList[T]) RemoveAtFront() {
	if s.head == nil {
		return
	}
	s.head = s.head.Next
	if s.head == nil {
		s.tail = nil
	}
	s.size--
}

func (s *SinglyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &SingleLinkNode[T]{Data: val}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		if s.tail != nil {
			s.tail.Next = newNode
		}
		s.tail = newNode
	}
	s.size++
}

func (s *SinglyLinkedList[T]) RemoveAtEnd() {
	if s.head == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
		s.size--
		return
	}
	currNode := s.head
	for currNode.Next != s.tail {
		currNode = currNode.Next
	}
	currNode.Next = nil
	s.tail = currNode
	s.size--
}

func (s *SinglyLinkedList[T]) Head() *SingleLinkNode[T] {
	return s.head
}

func (s *SinglyLinkedList[T]) Tail() *SingleLinkNode[T] {
	return s.tail
}

func (s *SinglyLinkedList[T]) Empty() bool {
	return s.head == nil
}

func (s *SinglyLinkedList[T]) Size() int {
	return s.size // size type is not a pointer
}
