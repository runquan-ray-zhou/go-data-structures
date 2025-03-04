package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*--------------------------------------------------------------------------------------------------*/
/* Test for NewSinglyLinkedList */
/*--------------------------------------------------------------------------------------------------*/
func TestNewSinglyLinkedList(t *testing.T) {

	// Happy Path
	t.Run("New singly linked list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
		assert.Equal(t, 0, list.size)
		assert.True(t, list.Empty())
	})

	// Edge Case
	t.Run("New singly linked list with different types", func(t *testing.T) {
		types := []any{
			NewSinglyLinkedList[int](),
			NewSinglyLinkedList[bool](),
			NewSinglyLinkedList[string](),
			NewSinglyLinkedList[[]int](),
			NewSinglyLinkedList[map[string]string](),
		}

		for _, tt := range types {
			assert.NotNil(t, tt)
		}
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for InsertAfter */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_InsertAfter(t *testing.T) {

	// Edge Case
	t.Run("Insert after in an empty list does nothing", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAfter(20, list.head)

		assert.Equal(t, 0, list.Size(), "List size should remain 0 when inserting after nil")
		assert.Nil(t, list.head, "Head should remain nil")
		assert.Nil(t, list.tail, "Tail should remain nil")
	})

	// Happy Path
	t.Run("Insert after head in single node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)

		assert.Equal(t, 20, list.head.Next.Data)
		assert.Equal(t, list.head.Next.Data, list.tail.Data)
		assert.Equal(t, list.Size(), 2)
	})

	// Happy Path
	t.Run("Insert after middle node", func(t *testing.T) {

		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)
		list.InsertAfter(30, list.head.Next)

		middleNode := list.head.Next

		list.InsertAfter(25, middleNode)

		assert.Equal(t, 10, list.head.Data)
		assert.Equal(t, 20, list.head.Next.Data)
		assert.Equal(t, 25, list.head.Next.Next.Data)
		assert.Equal(t, 30, list.head.Next.Next.Next.Data)
		assert.Equal(t, list.Size(), 4)

	})

	// Happy Path
	t.Run("Insert after tail in multi-node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)

		oldTail := list.tail

		list.InsertAfter(30, list.tail)

		assert.Equal(t, 30, list.tail.Data)
		assert.Equal(t, list.Size(), 3)
		assert.Same(t, oldTail.Next, list.tail)
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, list.head.Next.Next, list.tail)
	})

	// Empty/Null Case
	t.Run("Insert after nil node in non-empty list", func(t *testing.T) {

		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)

		originalSize := list.Size()
		originalHead := list.head
		originalTail := list.tail

		list.InsertAfter(30, nil)

		assert.Equal(t, list.Size(), originalSize)
		assert.Equal(t, list.head, originalHead)
		assert.Equal(t, list.tail, originalTail)
	})

	// Edge Case
	t.Run("Insert after node not in list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		foreignNode := &SingleLinkNode[int]{Data: 99}

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)

		originalSize := list.Size()
		originalHead := list.head
		originalTail := list.tail

		list.InsertAfter(30, foreignNode)

		assert.Equal(t, list.Size(), originalSize)
		assert.Same(t, list.head, originalHead)
		assert.Same(t, list.tail, originalTail)
	})

	// Data integrity
	t.Run("Data integrity through list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		// Create list: 10 -> 20 -> 30
		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)
		list.InsertAfter(30, list.head.Next)

		// Verify complete list structure
		current := list.head
		expected := []int{10, 20, 30}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for RemoveAfter */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_RemoveAfter(t *testing.T) {

	// Edge Case
	t.Run("Remove after from an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.RemoveAfter(list.head)

		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
		assert.Equal(t, 0, list.Size())
	})

	// Happy Path
	t.Run("Remove after head in single node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)

		list.RemoveAfter(list.head)

		assert.Equal(t, 10, list.head.Data)
		assert.Equal(t, list.tail, list.head)
		assert.Equal(t, 1, list.Size())
	})

	// Happy Path
	t.Run("Remove after middle node", func(t *testing.T) {

		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)
		list.InsertAfter(30, list.head.Next)

		middleNode := list.head.Next

		list.RemoveAfter(middleNode)

		assert.Equal(t, 10, list.head.Data)
		assert.Equal(t, 20, list.head.Next.Data)
		assert.Equal(t, 20, list.tail.Data)
		assert.Equal(t, list.Size(), 2)
	})

	// Edge Case
	t.Run("Remove after tail in multi-node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)

		list.RemoveAfter(list.head.Next)

		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 10, list.head.Data)
		assert.Equal(t, 20, list.head.Next.Data)
		assert.Equal(t, 2, list.Size())
	})

	// Empty/Null Cases
	t.Run("Remove after nil node in non-empty list", func(t *testing.T) {

		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)

		originalHead := list.head
		originalTail := list.tail

		list.RemoveAfter(list.tail)

		assert.Equal(t, 1, list.Size())
		assert.Equal(t, list.head, originalHead)
		assert.Equal(t, list.tail, originalTail)

	})

	// Data integrity
	t.Run("Data integrity through list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		// Create list: 10 -> 20 -> 30
		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)
		list.InsertAfter(30, list.head.Next)

		list.RemoveAfter(list.head.Next)

		// Verify complete list structure
		current := list.head
		expected := []int{10, 20}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})

	// Happy Path
	t.Run("Multiple consecutive removes", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)
		list.InsertAfter(30, list.head.Next)
		list.InsertAfter(40, list.head.Next.Next)

		list.RemoveAfter(list.head)
		list.RemoveAfter(list.head)

		assert.Equal(t, 10, list.head.Data)
		assert.Equal(t, 40, list.tail.Data)
		assert.Equal(t, 2, list.Size())
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for InsertAtFront */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_InsertAtFront(t *testing.T) {

	// Edge Case
	t.Run("Insert at front in an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)

		assert.NotNil(t, list.head)
		assert.NotNil(t, list.tail)
		assert.Equal(t, 1, list.Size())
		assert.Equal(t, list.head, list.tail)
		assert.Nil(t, list.tail.Next)
		assert.Nil(t, list.head.Next)
		assert.Equal(t, 10, list.head.Data)
	})

	// Happy Path
	t.Run("Insert at front in single node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(10)
		list.InsertAtFront(5)

		assert.Equal(t, 5, list.head.Data)
		assert.NotNil(t, list.head.Next)
		assert.Equal(t, 10, list.head.Next.Data)
		assert.Equal(t, list.tail, list.head.Next)
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 2, list.Size())
	})

	t.Run("Insert at front multiple nodes", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(5)
		list.InsertAtFront(4)
		list.InsertAtFront(3)
		list.InsertAtFront(2)
		list.InsertAtFront(1)

		assert.Equal(t, 5, list.Size())
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 5, list.tail.Data)
		assert.Equal(t, 1, list.head.Data)

		current := list.head
		expected := []int{1, 2, 3, 4, 5}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})

	// Data integrity
	t.Run("Data integrity through list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAfter(20, list.head)
		list.InsertAfter(30, list.head.Next)

		assert.Equal(t, 3, list.Size())
		assert.Equal(t, 30, list.tail.Data)
		assert.Nil(t, list.tail.Next)

		current := list.head
		expected := []int{10, 20, 30}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for RemoveAtFront */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_RemoveAtFront(t *testing.T) {

	// Edge Case
	t.Run("Remove at front in an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.RemoveAtFront()

		assert.Equal(t, 0, list.Size())
		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)
	})

	// Happy Path
	t.Run("Remove at front in single node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(10)
		list.RemoveAtFront()

		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.Empty())
	})

	t.Run("Remove at front multiple nodes", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(5)
		list.InsertAtFront(4)
		list.InsertAtFront(3)
		list.InsertAtFront(2)
		list.InsertAtFront(1)
		list.RemoveAtFront()
		list.RemoveAtFront()

		assert.Equal(t, 3, list.Size())
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 5, list.tail.Data)
		assert.Equal(t, 3, list.head.Data)

		current := list.head
		expected := []int{3, 4, 5}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})

	t.Run("Remove at front until empty", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		list.InsertAtEnd(2)
		list.InsertAtEnd(3)

		list.RemoveAtFront()
		list.RemoveAtFront()
		list.RemoveAtFront()

		assert.True(t, list.Empty())
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
		assert.Equal(t, 0, list.Size())
	})

	t.Run("Remove at front and then insert", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		list.InsertAtEnd(2)
		list.RemoveAtFront()
		list.InsertAtFront(3)

		assert.Equal(t, 2, list.Size())
		assert.Equal(t, 3, list.head.Data)
		assert.Equal(t, 2, list.tail.Data)
		assert.Nil(t, list.tail.Next)
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for InsertAtEnd */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_InsertAtEnd(t *testing.T) {

	// Edge Case
	t.Run("Insert at end in an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(10)

		assert.NotNil(t, list.head)
		assert.NotNil(t, list.tail)
		assert.Equal(t, 1, list.Size())
		assert.Equal(t, list.head, list.tail)
		assert.Nil(t, list.tail.Next)
		assert.Nil(t, list.head.Next)
		assert.Equal(t, 10, list.head.Data)
	})

	// Happy Path
	t.Run("Insert at end in single node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAtEnd(5)

		assert.Equal(t, 10, list.head.Data)
		assert.NotNil(t, list.head.Next)
		assert.Equal(t, 5, list.head.Next.Data)
		assert.Equal(t, list.tail, list.head.Next)
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 2, list.Size())
	})

	t.Run("Insert at end multiple nodes", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(5)
		list.InsertAtEnd(4)
		list.InsertAtEnd(3)
		list.InsertAtEnd(2)
		list.InsertAtEnd(1)

		assert.Equal(t, 5, list.Size())
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 1, list.tail.Data)
		assert.Equal(t, 5, list.head.Data)

		current := list.head
		expected := []int{5, 4, 3, 2, 1}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})

	// Data integrity
	t.Run("Data integrity through list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(10)
		list.InsertAtEnd(20)
		list.InsertAtEnd(30)

		assert.Equal(t, 3, list.Size())
		assert.Equal(t, 30, list.tail.Data)
		assert.Nil(t, list.tail.Next)

		current := list.head
		expected := []int{10, 20, 30}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for RemoveAtEnd */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_RemoveAtEnd(t *testing.T) {

	t.Run("Remove at end in an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.RemoveAtEnd()

		assert.Equal(t, 0, list.Size())
		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)
	})

	// Happy Path
	t.Run("Remove at end in single node list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.RemoveAtEnd()

		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.Empty())
	})

	t.Run("Remove at end front multiple nodes", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(5)
		list.InsertAtFront(4)
		list.InsertAtFront(3)
		list.InsertAtFront(2)
		list.InsertAtFront(1)
		list.RemoveAtEnd()
		list.RemoveAtEnd()
		list.RemoveAtEnd()

		assert.Equal(t, 2, list.Size())
		assert.Nil(t, list.tail.Next)
		assert.Equal(t, 2, list.tail.Data)
		assert.Equal(t, 1, list.head.Data)

		current := list.head
		expected := []int{1, 2}
		for i := 0; current != nil; i++ {
			assert.Equal(t, expected[i], current.Data)
			current = current.Next
		}
	})

	t.Run("Remove at end until empty", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		list.InsertAtEnd(2)
		list.InsertAtEnd(3)

		list.RemoveAtEnd()
		list.RemoveAtEnd()
		list.RemoveAtEnd()

		assert.True(t, list.Empty())
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
		assert.Equal(t, 0, list.Size())
	})

	t.Run("Remove at end and then insert", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		list.InsertAtEnd(2)
		list.RemoveAtEnd()
		list.InsertAtFront(3)

		assert.Equal(t, 2, list.Size())
		assert.Equal(t, 3, list.head.Data)
		assert.Equal(t, 1, list.tail.Data)
		assert.Nil(t, list.tail.Next)
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for Head */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_Head(t *testing.T) {

	t.Run("Return nil when head is empty", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		listHead := list.Head()

		assert.Nil(t, listHead)
	})

	t.Run("Return head of single node linked list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		listHead := list.Head()
		listHeadData := listHead.Data

		assert.Equal(t, 1, listHeadData)
		assert.Equal(t, 1, list.Size())
		assert.Equal(t, list.head, listHead)
		assert.Nil(t, listHead.Next)

	})

	t.Run("Return head after operations", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		list.InsertAtFront(100)
		list.RemoveAtFront()
		list.InsertAtFront(50)
		listHead := list.Head()
		listHeadData := listHead.Data

		assert.Equal(t, 50, listHeadData)
		assert.NotNil(t, listHead)
		assert.Equal(t, 2, list.Size())
		assert.Equal(t, 1, listHead.Next.Data)
		assert.Nil(t, listHead.Next.Next)
	})

}

/*--------------------------------------------------------------------------------------------------*/
/* Test for Tail */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_Tail(t *testing.T) {

	t.Run("Return nil when tail is nil for an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		listTail := list.Tail()

		assert.Nil(t, listTail)
	})

	t.Run("Return tail of single node linked list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		listTail := list.Tail()
		listTailData := listTail.Data

		assert.Equal(t, 1, listTailData)
		assert.Equal(t, 1, list.Size())
		assert.Equal(t, list.tail, listTail)
		assert.Nil(t, listTail.Next)

	})

	t.Run("Return tail after operations", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtEnd(1)
		list.InsertAtFront(100)
		list.RemoveAtFront()
		list.InsertAtFront(50)
		listTail := list.Tail()
		listTailData := listTail.Data

		assert.Equal(t, 1, listTailData)
		assert.NotNil(t, listTail)
		assert.Equal(t, 2, list.Size())
	})
}

/*--------------------------------------------------------------------------------------------------*/
/* Test for Empty */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_Empty(t *testing.T) {

	t.Run("Return Empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.RemoveAtEnd()

		assert.Empty(t, list)
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
	})

	t.Run("Return true for empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.RemoveAtEnd()
		isEmpty := list.Empty()

		assert.True(t, isEmpty)
		assert.Empty(t, list)
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
	})

	t.Run("Return false for non empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.RemoveAtEnd()
		list.InsertAtEnd(20)
		isEmpty := list.Empty()

		assert.NotEmpty(t, list)
		assert.Equal(t, 1, list.size)
		assert.False(t, isEmpty)
		assert.NotNil(t, list.head)
		assert.Equal(t, 20, list.head.Data)
	})

}

/*--------------------------------------------------------------------------------------------------*/
/* Test for Size */
/*--------------------------------------------------------------------------------------------------*/

func TestSinglyLinkedList_Size(t *testing.T) {

	t.Run("Return correct size", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAtEnd(20)
		listSize := list.Size()

		assert.Equal(t, 2, listSize)
		assert.NotEmpty(t, list)
		assert.NotNil(t, list.tail)
		assert.Equal(t, 20, list.tail.Data)
	})

	t.Run("Return size zero for empty list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAtEnd(20)
		list.RemoveAtFront()
		list.RemoveAtEnd()
		listSize := list.Size()

		assert.Equal(t, 0, listSize)
		assert.Empty(t, list)
		assert.Nil(t, list.head)
		assert.Empty(t, list.tail)
	})

	t.Run("Size updates correctly after removals", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		list.InsertAtFront(10)
		list.InsertAtEnd(20)
		listSizeBefore := list.Size()
		list.RemoveAtFront()
		list.RemoveAtEnd()
		listSizeAfter := list.Size()

		assert.Equal(t, 2, listSizeBefore)
		assert.Equal(t, 0, listSizeAfter)
		assert.Equal(t, 0, list.size)
	})
}
