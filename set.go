package godatastructures

import (
	"encoding/json"
	"golang.org/x/exp/constraints"
	"hash/fnv"
)

type SetInterface[T constraints.Ordered] interface {
	hash(val T) int      // helper function to assign val to an index in the underlying array
	resize()             // helper function to double the size of the underlying array
	Contains(val T) bool // returns whether set contains element val. O(1)
	Insert(val T)        // inserts val in set if not present, increases size by 1. O(1)
	Remove(val T)        // removes item from set if present, decreases size by 1. O(1)
	Empty() bool         // returns whether set is empty. O(1)
	Size() int           // returns number of elements in set. O(1)
	Values() []T         // returns all values in the set.
}

// Set should implement SetInterface; write a constructor & methods to complete it
type Set[T constraints.Ordered] struct {
	arr     []BinarySearchTree[T]
	maxFill float32 // value in (0, 1] indicating the maximum allowable ratio of elements to arr before arr is resized
	size    int
}
