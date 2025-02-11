package graph

import (
	"golang.org/x/exp/constraints"
)

type GraphInterface[T constraints.Ordered] interface {
	Insert(val Pair[T], neighbors []T) // helper function to build/update graph. inserts T into graph if not present, adds neighbors. O(neighbors)
	Remove(val T)                      // helper function to build/update graph. removes T from graph, removes T from other nodes' neighbors.  O(neighbors)
	Empty() bool                       // returns whether graph is empty. O(1)
	DepthFirstTraversal() []T          // returns values in graph ordered by DFS processing. O(n)
	BreadthFirstTraversal() []T        // returns values in graph ordered by BFS processing. O(n)
	Size() int                         // returns number of items in graph.
}

type Graph[T constraints.Ordered] struct {
	nodes map[T]struct{}
	// Using second map as a set; ignore the value
	// Ensure on removals you use delete func
	neighbors map[T]map[T]struct{}
}

func NewGraph[T constraints.Ordered]() *Graph[T] {
	return &Graph[T]{
		nodes:     make(map[T]struct{}),
		neighbors: make(map[T]map[T]struct{}),
	}
}

func (g *Graph[T]) Empty() bool {
	return len(g.nodes) == 0
}

func (g *Graph[T]) Size() int {
	return len(g.nodes)
}

func (g *Graph[T]) Insert(pair Pair[T], neighbors []T) {
	key := pair.Key

	g.nodes[key] = struct{}{}

	if g.neighbors[key] == nil {
		g.neighbors[key] = make(map[T]struct{})
	}

	for _, neighbor := range neighbors {
		g.nodes[neighbor] = struct{}{}
		if g.neighbors[neighbor] == nil {
			g.neighbors[neighbor] = make(map[T]struct{})
		}
		g.neighbors[key][neighbor] = struct{}{} // empty struct type have no fields and takes up zero bytes of memory
		g.neighbors[neighbor][key] = struct{}{} // great to use to track existence of key without any associated data
	}

	// set := make(map[T]struct{})
	// for _, neighbor := range neighbors {
	// 	set[neighbor] = struct{}{}
	// }
	// g.neighbors[key] = set
	// for _, neighbor := range neighbors {
	// 	g.neighbors[neighbor][key] = struct{}{}
	// }
}

func (g *Graph[T]) Remove(key T) {
	delete(g.nodes, key)
	neighbors := g.neighbors[key]
	for neighbor, _ := range neighbors {
		delete(g.neighbors[neighbor], key)
	}
	delete(g.neighbors, key)
}

// did not give starting node.
func (g *Graph[T]) DepthFirstTraversal() []T {
	visited := make(map[T]struct{})
	result := []T{}

	var helperDFS func(node T)
	helperDFS = func(node T) {
		if _, ok := visited[node]; ok {
			return
		}

		visited[node] = struct{}{}
		result = append(result, node)

		for neighbor := range g.neighbors[node] {
			helperDFS(neighbor)
		}
	}

	for node := range g.nodes {
		if _, ok := visited[node]; !ok {
			helperDFS(node)
		}
	}

	return result
}

//sorted version
// func (g *Graph[T]) DepthFirstTraversal() []T {
// 	visited := make(map[T]struct{})
// 	result := []T{}

// 	var helperDFS func(node T)
// 	helperDFS = func(node T) {
// 		if _, ok := visited[node]; ok {
// 			return
// 		}

// 		visited[node] = struct{}{}
// 		result = append(result, node)

// 		neighbors := []T{}
// 		for neighbor := range g.neighbors[node] {
// 			neighbors = append(neighbors, neighbor)
// 		}

// 		sort.Slice(neighbors, func(i, j int) bool {
// 			return neighbors[i] < neighbors[j]
// 		})

// 		for _, neighbor := range neighbors {
// 			helperDFS(neighbor)
// 		}
// 	}

// 	nodes := []T{}
// 	for node := range g.nodes {
// 		nodes = append(nodes, node)
// 	}

// 	sort.Slice(nodes, func(i, j int) bool {
// 		return nodes[i] < nodes[j]
// 	})

// 	for node := range g.nodes {
// 		if _, ok := visited[node]; !ok {
// 			helperDFS(node)
// 		}
// 	}
// 	return result
// }

// dont use recursion
func (g *Graph[T]) BreadthFirstTraversal() []T {
	visited := make(map[T]struct{})
	result := []T{}
	queue := NewQueue[T]()

	for node := range g.nodes {
		if _, ok := visited[node]; ok {
			continue
		}

		queue.Enqueue(node)

		for queue.Size() > 0 {
			front := queue.Front()
			queue.Dequeue()

			if _, ok := visited[front]; ok {
				continue
			}

			visited[front] = struct{}{}
			result = append(result, front)

			for neighbor := range g.neighbors[front] {
				if _, ok := visited[neighbor]; !ok {
					queue.Enqueue(neighbor)
				}
			}
		}
	}
	return result
}

/*
undirected - edges, source and destination
n nodes, max possible number of edges
n square
sparse vs dense. the number of edges vs the number of nodes.
O(N + E) big o notation
weighted vs weighted
edges represented by time or distance/strength of connection.
graph algorithm, distance between nodes.
weighted vs unweighted
does a path exits problem
shortest path
batching orders
what is the shortest path to pick up the food.  outsourcing mapping.  uber does the matching part.

unweighted class is the number of edges.

weighted is the sum of the edges.

what is the intuitive way to be represented in a data

every earner is a node
every city is node

adjacency list/set notation, set is use more often
map "node"(unique id) -> list/set of "nodes"(unique ids)
NYC -> (New haven, hartford, baltimore, dc, boston)
new haven -> (new york , hartford)
bermuda -> ()

weighted version, store the edge to the node


adjacency matrix

recursive branches

if there is a single cycle the graph will be acyclic.

searching for something
dfs, we going to see something very far away before we see something closer


mechanics of traversing a graph
use the information from edges choose where to go
greedy heuristic algorithm

use the weight of the edges to traverse in a smart
use a matrix

to populate

algorithm

Dijkstra's algorithm
A* algorithm
Bellman–Ford algorithm
Floyd–Warshall algorithm

graphs
how many sources and destination

anary graph

*/
