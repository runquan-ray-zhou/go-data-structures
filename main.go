package main

import "fmt"

func main() {
	node := &SingleLinkNode[int]{Data: 3}
	fmt.Print(node)
	list := SinglyLinkedList[int]{head: node}
	list.InsertAfter(4, node)
	fmt.Print(list.head.Data)
}
