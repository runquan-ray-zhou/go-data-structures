package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSinglyLinkedList(t *testing.T) {

	t.Run("New singly linked list is empty", func(t *testing.T) {
		expectedList := NewSinglyLinkedList[int]()

		assert.Nil(t, expectedList.head)
	})
}
