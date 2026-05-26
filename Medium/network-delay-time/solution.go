package main

import (
	"container/heap"
	"fmt"
	"math"
)

// heap node
type Item struct {
	node int
	dist int
}

// minheap implementation
type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func networkDelayTime(times [][]int, n int, k int) int {
	// build adjacency list
	graph := make(map[int][]Item)

	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		graph[u] = append(graph[u], Item{
			node: v,
			dist: w,
		})
	}

	// distance array
	dist := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}

	dist[k] = 0

	// min heap
	h := &MinHeap{}
	heap.Init(h)

	heap.Push(h, Item{
		node: k,
		dist: 0,
	})

	// dijkstra's algorithm
	for h.Len() > 0 {
		curr := heap.Pop(h).(Item)

		node := curr.node
		currDist := curr.dist

		// skip outdated entries
		if currDist > dist[node] {
			continue
		}

		// explore neighbors
		for _, nei := range graph[node] {
			nextNode := nei.node
			newDist := currDist + nei.dist

			if newDist < dist[nextNode] {
				dist[nextNode] = newDist

				heap.Push(h, Item{
					node: nextNode,
					dist: newDist,
				})
			}
		}
	}

	// find maximum shortest distance
	maxTime := 0

	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}

		if dist[i] > maxTime {
			maxTime = dist[i]
		}
	}

	return maxTime
}

func main() {
	times := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}

	fmt.Println(networkDelayTime(times, 4, 2)) // output: 2
}
