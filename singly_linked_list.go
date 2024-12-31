package godatastructures

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
	// size int
}
