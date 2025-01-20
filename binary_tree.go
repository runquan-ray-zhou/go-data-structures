package main

import (
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct { //Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >. If future releases of Go add new ordered types, this constraint will be modified to include them.
	Key   T
	Value any
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

type BinaryTreeInterface[T constraints.Ordered] interface {
	Root() *BinaryTreeNode[T]       // returns node at root of tree. O(1)
	Insert(key T, val any)          // inserts node with key, val (can be nil), increases size by 1. O(logn)
	Contains(key T) bool            // checks if tree contains key. O(logn)
	Lookup(key T) any               // returns value for key. O(logn)
	Update(key T, val any)          // update key with new value O(logn)
	Remove(key T)                   // removes node with key, decreases size by 1. O(logn)
	Empty() bool                    // returns whether tree is empty. O(1)
	InOrderTraversal() []Pair[T]    // returns keys in tree ordered by processing left, current, right. O(n)
	PreOrderTraversal() []Pair[T]   // returns keys in tree ordered by processing current, left, right. O(n)
	PostOrderTraversal() []Pair[T]  // returns keys in tree ordered by processing left, right, current. O(n)
	LevelOrderTraversal() []Pair[T] // returns keys in tree ordered by processing each level left to right. O(n)
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{nil}
}

func (bst *BinarySearchTree[T]) Root() *BinaryTreeNode[T] {
	return bst.root
}

func (bst *BinarySearchTree[T]) Insert(key T, val any) {
	newNode := &BinaryTreeNode[T]{key, val, nil, nil}
	if bst.root == nil {
		bst.root = newNode
		return
	}
	curr := bst.root
	for curr != nil {
		if key < curr.Key {
			if curr.Left == nil {
				curr.Left = newNode
				return
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = newNode
				return
			}
			curr = curr.Right
		}
	}
}

func (bst *BinarySearchTree[T]) Contains(key T) bool {
	curr := bst.root
	for curr != nil {
		if key == curr.Key {
			return true
		} else if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return false
}

func (bst *BinarySearchTree[T]) Lookup(key T) any {
	curr := bst.root
	for curr != nil {
		if key == curr.Key {
			return curr.Value
		} else if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return nil
}

func (bst *BinarySearchTree[T]) Update(key T, val any) {
	curr := bst.root
	for curr.Key != key {
		if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	curr.Value = val
}

func (bst *BinarySearchTree[T]) Remove(key T) {
	parent := (*BinaryTreeNode[T])(nil)
	curr := bst.root
	// Find node & parent
	for curr.Key != key {
		parent = curr
		if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	if curr.Left == nil && curr.Right == nil { // Node has no children
		if curr == bst.root { // Node is root
			bst.root = nil
		} else if parent.Left == curr { // Node is left child
			parent.Left = nil
		} else { // Node is right child
			parent.Right = nil
		}
	} else if curr.Left == nil { // Node has right child
		if curr == bst.root { // Node is root
			bst.root = curr.Right
		} else if parent.Left == curr { // Node is left child
			parent.Left = curr.Right
		} else { // Node is right child
			parent.Right = curr.Right
		}
	} else if curr.Right == nil { // Node has left child
		if curr == bst.root { // Node is root
			bst.root = curr.Left
		} else if parent.Left == curr { // Node is left child
			parent.Left = curr.Left
		} else { // Node is right child
			parent.Right = curr.Left
		}
	} else { // Node has both children
		next := curr.Right
		for next.Left != nil {
			next = next.Left
		}
		newKey := next.Key
		newVal := next.Value
		bst.Remove(newKey)
		curr.Key = newKey
		curr.Value = newVal
	}
}

func (bst *BinarySearchTree[T]) Empty() bool {
	return bst.root == nil
}

func (bst *BinarySearchTree[T]) InOrderTraversal() []Pair[T] {
	// pairs := []Pair[T]{}
	// if bst.root == nil {
	// 	return pairs
	// }
	// leftTree := &BinarySearchTree[T]{root: bst.root.Left}
	// rightTree := &BinarySearchTree[T]{root: bst.root.Right}
	// curr := bst.root
	// if curr.Left != nil {
	// 	pairs = append(pairs, leftTree.InOrderTraversal()...)
	// }
	// currPair := Pair[T]{Key: curr.Key, Value: curr.Value}
	// pairs = append(pairs, currPair)

	// if curr.Right != nil {
	// 	pairs = append(pairs, rightTree.InOrderTraversal()...)
	// }
	// return pairs
	// Below use helper function for more efficiency
	pairs := []Pair[T]{}
	var inOrder func(node *BinaryTreeNode[T])

	inOrder = func(node *BinaryTreeNode[T]) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		pairs = append(pairs, Pair[T]{Key: node.Key, Value: node.Value})
		inOrder(node.Right)
	}

	inOrder(bst.root)
	return pairs
}

func (bst *BinarySearchTree[T]) PreOrderTraversal() []Pair[T] {
	pairs := []Pair[T]{}
	var preOrder func(node *BinaryTreeNode[T])

	preOrder = func(node *BinaryTreeNode[T]) {
		if node == nil {
			return
		}
		pairs = append(pairs, Pair[T]{Key: node.Key, Value: node.Value})
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(bst.root)
	return pairs
}

func (bst *BinarySearchTree[T]) PostOrderTraversal() []Pair[T] {
	pairs := []Pair[T]{}
	var postOrder func(node *BinaryTreeNode[T])

	postOrder = func(node *BinaryTreeNode[T]) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		pairs = append(pairs, Pair[T]{Key: node.Key, Value: node.Value})
	}

	postOrder(bst.root)
	return pairs
}

func (bst *BinarySearchTree[T]) LevelOrderTraversal() []Pair[T] {
	pairs := []Pair[T]{}

	if bst.root == nil {
		return pairs
	}

	// A queue to keep track of nodes to process
	queue := []*BinaryTreeNode[T]{bst.root}

	for len(queue) > 0 {
		// Dequeue the first node from the queue
		curr := queue[0]
		queue = queue[1:]

		// Process the current node
		pairs = append(pairs, Pair[T]{Key: curr.Key, Value: curr.Value})

		// Enqueue the left child if it exists
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

		// Enqueue the right child if it exists
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return pairs
}
