package main

import "fmt"

func main() {
	node := &SingleLinkNode[int]{Data: 3}
	fmt.Println(node)
	list := SinglyLinkedList[int]{head: node, tail: nil, size: 1}
	fmt.Println(list.size)
	list.InsertAfter(4, node)
	fmt.Println(list.size)
	fmt.Println(list.head.Data)
	stack := &Stack[int]{list: &list}
	stack.Pop()
	fmt.Println(list.head.Data)
	fmt.Println(list.size)
}
