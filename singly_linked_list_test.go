package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSinglyLinkedList(t *testing.T) {

	t.Run("New singly linked list is empty", func(t *testing.T) {
		expectedList := NewSinglyLinkedList[int]()

		assert.Nil(t, expectedList.head, "Head should remain nil")
		assert.Nil(t, expectedList.tail, "Tail should remain nil")
		assert.Empty(t, expectedList.Empty(), "List should remain empty")

	})
}

func TestSinglyLinkedListInsertAtFront(t *testing.T) {
	t.Run("Insert at front", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtFront(2)

		assert.Equal(t, 2, list.head.Data, "Head should be 2")
		assert.Equal(t, 1, list.tail.Data, "Tail should be 1")
		assert.Equal(t, 2, list.Size(), "Size should be 2")
	})
}

func TestSinglyLinkedListInsertAfter(t *testing.T) {
	t.Run("Insert after", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAfter(2, list.head)

		assert.Equal(t, 1, list.head.Data, "Head should be 1")
		assert.Equal(t, 2, list.head.Next.Data, "Head next should be 2")
		assert.Equal(t, 2, list.tail.Data, "Tail should be 2")
		assert.Equal(t, 2, list.Size(), "Size should be 2")
	})
}

func TestSinglyLinkedListInsertAtEnd(t *testing.T) {
	t.Run("Insert at end", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtEnd(2)

		assert.Equal(t, 1, list.head.Data, "Head should be 1")
		assert.Equal(t, 2, list.tail.Data, "Tail should be 2")
		assert.Equal(t, 2, list.Size(), "Size should be 2")
	})
}

func TestSinglyLinkedListRemoveAtEnd(t *testing.T) {
	t.Run("Remove at end", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtEnd(2)
		list.RemoveAtEnd()

		assert.Equal(t, 1, list.head.Data, "Head should be 1")
		assert.Equal(t, 1, list.tail.Data, "Tail should be 1")
		assert.Equal(t, 1, list.Size(), "Size should be 1")
	})
}

func TestSinglyLinkedListRemoveAtFront(t *testing.T) {
	t.Run("Remove at front", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtFront(2)
		list.RemoveAtFront()

		assert.Equal(t, 1, list.head.Data, "Head should be 1")
		assert.Equal(t, 1, list.tail.Data, "Tail should be 1")
		assert.Equal(t, 1, list.Size(), "Size should be 1")
	})
}

func TestSinglyLinkedListRemoveAfter(t *testing.T) {
	t.Run("Remove after", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAfter(2, list.head)
		list.RemoveAfter(list.head)

		assert.Equal(t, 1, list.head.Data, "Head should be 1")
		assert.Equal(t, 1, list.tail.Data, "Tail should be 1")
		assert.Equal(t, 1, list.Size(), "Size should be 1")
	})
}

func TestSinglyLinkedListEmpty(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtFront(2)
		list.RemoveAtFront()

		assert.False(t, list.Empty(), "List should not be empty")
	})
}

func TestSinglyLinkedListSize(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtFront(2)
		list.RemoveAtFront()

		assert.Equal(t, 1, list.Size(), "Size should be 1")
	})
}

func TestSinglyLinkedListHead(t *testing.T) {
	t.Run("Head", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtFront(2)

		assert.Equal(t, 2, list.Head().Data, "Head should be 2")
	})
}

func TestSinglyLinkedListTail(t *testing.T) {
	t.Run("Tail", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.InsertAtFront(1)
		list.InsertAtFront(2)

		assert.Equal(t, 1, list.Tail().Data, "Tail should be 1")
	})
}
