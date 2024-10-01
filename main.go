package main

import (
	"fmt"
	"container/heap"
	"math"
)

// Item represents an item in the priority queue.
type Item struct {
	node     int     // The node number
	priority float64 // The priority (i.e., the current known shortest distance)
	index    int     // The index in the heap
}

// PriorityQueue implements a priority queue with heap.Interface.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // Avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Graph structure where each node points to its neighbors and their weights
type Graph map[int]map[int]float64

// Dijkstra finds the shortest paths from a source node to all other nodes in the graph.
func Dijkstra(graph Graph, start int) map[int]float64 {
	distances := make(map[int]float64)
	for node := range graph {
		distances[node] = math.Inf(1) // Initialize distances to infinity
	}
	distances[start] = 0

	// Priority queue for the next node to explore
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, priority: 0})

	// While there are nodes to explore
	for pq.Len() > 0 {
		// Get the node with the smallest known distance
		item := heap.Pop(&pq).(*Item)
		currentNode := item.node
		currentDistance := item.priority

		// Process each neighbor of the current node
		for neighbor, weight := range graph[currentNode] {
			distance := currentDistance + weight

			// If a shorter path to the neighbor is found
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(&pq, &Item{node: neighbor, priority: distance})
			}
		}
	}

	return distances
}

func main() {
	// Example usage of the Dijkstra function
	graph := Graph{
		0: {1: 4, 2: 1},
		1: {3: 1},
		2: {1: 2, 3: 5},
		3: {},
	}

	// Compute shortest paths from node 0
	start := 0
	distances := Dijkstra(graph, start)

	// Print the distances to each node from the start
	fmt.Printf("Shortest distances from node %d:\n", start)
	for node, distance := range distances {
		fmt.Printf("Node %d: %f\n", node, distance)
	}
}
