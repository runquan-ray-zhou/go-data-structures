package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// graph := NewGraph[string]()

	// graph.Insert(Pair[string]{Key: "Alice", Value: nil}, []string{"Bob", "Candy", "Derek", "Elaine"})
	// graph.Insert(Pair[string]{Key: "Bob", Value: nil}, []string{"Fred", "Alice"})
	// graph.Insert(Pair[string]{Key: "Fred", Value: nil}, []string{"Bob", "Helen"})
	// graph.Insert(Pair[string]{Key: "Helen", Value: nil}, []string{"Fred", "Candy"})
	// graph.Insert(Pair[string]{Key: "Candy", Value: nil}, []string{"Alice", "Helen"})
	// graph.Insert(Pair[string]{Key: "Derek", Value: nil}, []string{"Alice", "Elaine", "Gina"})
	// graph.Insert(Pair[string]{Key: "Gina", Value: nil}, []string{"Derek", "Irena"})
	// graph.Insert(Pair[string]{Key: "Irena", Value: nil}, []string{"Gina"})
	// graph.Insert(Pair[string]{Key: "Elaine", Value: nil}, []string{"Alice", "Derek"})

	/*
			graph.nodes
			map[string]struct{}{
				"Alice":  struct{}{},
				"Bob":    struct{}{},
				"Candy":  struct{}{},
				"Derek":  struct{}{},
				"Elaine": struct{}{},
				"Fred":   struct{}{},
				"Gina":   struct{}{},
				"Helen":  struct{}{},
				"Irena":  struct{}{},
			}

			graph.neighbors
			map[string]map[string]struct{}{
		    "Alice": {
		        "Bob":    struct{}{},
		        "Candy":  struct{}{},
		        "Derek":  struct{}{},
		        "Elaine": struct{}{},
		    },
		    "Bob": {
		        "Fred":   struct{}{},
		        "Alice":  struct{}{},
		    },
		    "Fred": {
		        "Bob":    struct{}{},
		        "Helen":  struct{}{},
		    },
		    "Helen": {
		        "Fred":   struct{}{},
		        "Candy":  struct{}{},
		    },
		    "Candy": {
		        "Alice":  struct{}{},
		        "Helen":  struct{}{},
		    },
		    "Derek": {
		        "Alice":  struct{}{},
		        "Elaine": struct{}{},
		        "Gina":   struct{}{},
		    },
		    "Elaine": {
		        "Alice":  struct{}{},
		        "Derek":  struct{}{},
		    },
		    "Gina": {
		        "Derek":  struct{}{},
		        "Irena":  struct{}{},
		    },
		    "Irena": {
		        "Gina":   struct{}{},
		    },
		}
	*/

	// fmt.Println("Graph structure:")
	// for node, neighbors := range graph.neighbors {
	// 	fmt.Printf("%s -> ", node)
	// 	for neighbor := range neighbors {
	// 		fmt.Printf("%s ", neighbor)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println(graph.Size())
	// fmt.Println(graph.Empty())

	// visited := graph.DepthFirstTraversal()

	// fmt.Println(visited)

	// result := graph.BreadthFirstTraversal()
	// fmt.Println(result)

	// node := &SingleLinkNode[int]{Data: 3}
	// fmt.Println(node)
	// list := SinglyLinkedList[int]{head: node, tail: nil, size: 1}
	// fmt.Println(list.size)
	// list.InsertAfter(4, node)
	// fmt.Println(list.size)
	// fmt.Println(list.head.Data)
	// stack := &Stack[int]{list: &list}
	// stack.Pop()
	// fmt.Println(list.head.Data)
	// fmt.Println(list.size)

	// circularArrayQueue := &AlternateQueue[string]{
	// 	arr:   []string{"a", "b", "c", "d", "e", "f", "g"},
	// 	first: 0,
	// 	last:  6,
	// }
	// fmt.Println(circularArrayQueue.Size())
	// fmt.Println(circularArrayQueue.Empty())
	// fmt.Println(circularArrayQueue.Front())
	// circularArrayQueue.Enqueue("h")
	// fmt.Println(circularArrayQueue.Size())
}
