package godatastructures

import (
	"encoding/json"
	"golang.org/x/exp/constraints"
	"hash/fnv"
)

type MapInterface[T constraints.Ordered, V any] interface {
	hash(key T) int      // helper function to assign val to an index in the underlying array
	resize()             // helper function to double the size of the underlying array
	Contains(key T) bool // returns whether map contains element key. O(1)
	Get(key T)           // returns value for key in map. throws error if not present. O(1)
	Set(key T, val V)    // inserts key-val pair in map if key is not present, increases size by 1. set key's value to val if present. O(1)
	Removes(key T)       // removes key from map if present, decreases size by 1. O(1)
	Empty() bool         // returns whether map is empty. O(1)
	Size() int           // returns number of elements in map. O(1)
	Values() []any       // returns all values in the map.
	Keys() []T           // returns all keys in the map.
	Objects() []Pair[T]  // returns all key-value pairs in the map.
}

// Map should implement MapInterface; write a constructor & methods to complete it
type Map[T constraints.Ordered, V any] struct {
	arr     []BinarySearchTree[T]
	maxFill float32 // value in (0, 1] indicating the maximum allowable ratio of elements to arr before arr is resized
	size    int
}
