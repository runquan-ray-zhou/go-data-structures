package doublylinkedlist

type DoubleLinkNode[T any] struct {
	Data T
	Next *DoubleLinkNode[T]
	Prev *DoubleLinkNode[T]
}

type DoublyLinkedListInterface[T any] interface {
	InsertAfter(val T, prev *DoubleLinkNode[T])  // create new node with data val after node prev, increase size by 1 O(1)
	InsertBefore(val T, next *DoubleLinkNode[T]) // create new node with data val before node next, increase size by 1 O(1)
	Remove(node *DoubleLinkNode[T])              // remove node , decrease size by 1 O(1)
	InsertAtFront(val T)                         // create node with data val at front of list, increase size by 1. O(1)
	RemoveAtFront()                              // remove node at front of list, decrease size by 1. O(1)
	InsertAtEnd(val T)                           // create node with data val at end of list, increase size by 1. O(n)
	RemoveAtEnd()                                // remove node at end of list, decrease size by 1. O(n)
	Head() *DoubleLinkNode[T]                    // return first node in list. O(1)
	Tail() *DoubleLinkNode[T]                    // return last node in list. O(1)
	Empty() bool                                 // returns whether list is empty. O(1)
	Size() int                                   // returns number of elements in list. O(1) or O(n) depending on implementation
}

type DoublyLinkedList[T any] struct {
	head *DoubleLinkNode[T]
	tail *DoubleLinkNode[T]
	// you can choose whether to store the size by uncommenting the following line
	size int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (s *DoublyLinkedList[T]) InsertAfter(val T, prev *DoubleLinkNode[T]) {
	if prev == nil { //Check for nil before insert
		s.InsertAtFront(val) //Adding head node
		return               //Prevent nil dereference
	}
	newNode := &DoubleLinkNode[T]{Data: val, Prev: prev, Next: prev.Next} // include newNode.Prev and newNode.Next in instantiation
	if prev.Next != nil {
		prev.Next.Prev = newNode
	} else {
		s.tail = newNode
	}
	prev.Next = newNode
	s.size++
}

func (s *DoublyLinkedList[T]) InsertBefore(val T, next *DoubleLinkNode[T]) {
	if next == nil { //Check for nil before insert
		s.InsertAtEnd(val) //Adding tail node
		return             //Prevent nil dereference
	}
	newNode := &DoubleLinkNode[T]{Data: val, Next: next, Prev: next.Prev}
	if next.Prev != nil {
		next.Prev.Next = newNode
	} else {
		s.head = newNode
	}
	next.Prev = newNode
	s.size++
}

func (s *DoublyLinkedList[T]) Remove(node *DoubleLinkNode[T]) {
	if s.head == nil || node == nil {
		return
	}
	if node == s.head {
		s.RemoveAtFront()
		return
	}
	if node == s.tail {
		s.RemoveAtEnd()
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	s.size--
}

func (s *DoublyLinkedList[T]) InsertAtFront(val T) {
	newNode := &DoubleLinkNode[T]{Data: val}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.head.Prev = newNode
		newNode.Next = s.head
		s.head = newNode
	}
	s.size++
}

func (s *DoublyLinkedList[T]) RemoveAtFront() {
	if s.head == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.head.Next.Prev = nil
		s.head = s.head.Next
	}
	s.size--
}

func (s *DoublyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &DoubleLinkNode[T]{Data: val}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.Next = newNode
		newNode.Prev = s.tail
		s.tail = newNode
	}
	s.size++
}

func (s *DoublyLinkedList[T]) RemoveAtEnd() {
	if s.head == nil || s.tail == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.tail = s.tail.Prev
		s.tail.Next = nil
	}
	s.size--
}

func (s *DoublyLinkedList[T]) Head() *DoubleLinkNode[T] {
	return s.head
}

func (s *DoublyLinkedList[T]) Tail() *DoubleLinkNode[T] {
	return s.tail
}

func (s *DoublyLinkedList[T]) Empty() bool {
	return s.head == nil
}

func (s *DoublyLinkedList[T]) Size() int {
	return s.size // size type is not a pointer
}
