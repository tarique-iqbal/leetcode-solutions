package main

import (
	"testing"
)

func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))

	for i := range adjList {
		nodes[i] = &Node{Val: i + 1}
	}

	for i, neighbors := range adjList {
		for _, nei := range neighbors {
			nodes[i].Neighbors = append(
				nodes[i].Neighbors,
				nodes[nei-1],
			)
		}
	}

	return nodes[0]
}

func deepEqualGraph(n1, n2 *Node) bool {
	return deepEqualHelper(n1, n2, make(map[*Node]*Node))
}

func deepEqualHelper(n1, n2 *Node, visited map[*Node]*Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if mapped, ok := visited[n1]; ok {
		return mapped == n2
	}

	if n1.Val != n2.Val {
		return false
	}

	if len(n1.Neighbors) != len(n2.Neighbors) {
		return false
	}

	visited[n1] = n2

	for i := range n1.Neighbors {
		if !deepEqualHelper(
			n1.Neighbors[i],
			n2.Neighbors[i],
			visited,
		) {
			return false
		}
	}

	return true
}

func verifyDeepCopy(original, cloned *Node) bool {
	if original == nil && cloned == nil {
		return true
	}

	if original == nil || cloned == nil {
		return false
	}

	type pair struct {
		orig  *Node
		clone *Node
	}

	queue := []pair{
		{orig: original, clone: cloned},
	}

	visited := make(map[*Node]*Node)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		orig := current.orig
		clone := current.clone

		// Must not be same pointer
		if orig == clone {
			return false
		}

		if mapped, ok := visited[orig]; ok {
			if mapped != clone {
				return false
			}
			continue
		}

		visited[orig] = clone

		if orig.Val != clone.Val {
			return false
		}

		if len(orig.Neighbors) != len(clone.Neighbors) {
			return false
		}

		for i := range orig.Neighbors {
			queue = append(queue, pair{
				orig:  orig.Neighbors[i],
				clone: clone.Neighbors[i],
			})
		}
	}

	return true
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "Empty graph",
			adjList: [][]int{},
		},
		{
			name: "Single node",
			adjList: [][]int{
				{},
			},
		},
		{
			name: "Single node self loop",
			adjList: [][]int{
				{1},
			},
		},
		{
			name: "Two node cycle",
			adjList: [][]int{
				{2},
				{1},
			},
		},
		{
			name: "LeetCode example",
			adjList: [][]int{
				{2, 4},
				{1, 3},
				{2, 4},
				{1, 3},
			},
		},
		{
			name: "Tree structure",
			adjList: [][]int{
				{2, 3},
				{},
				{},
			},
		},
		{
			name: "Long chain",
			adjList: [][]int{
				{2},
				{3},
				{4},
				{},
			},
		},
		{
			name: "Triangle cycle",
			adjList: [][]int{
				{2, 3},
				{1, 3},
				{1, 2},
			},
		},
		{
			name: "Diamond graph",
			adjList: [][]int{
				{2, 3},
				{1, 4},
				{1, 4},
				{2, 3},
			},
		},
		{
			name: "Complete graph of four nodes",
			adjList: [][]int{
				{2, 3, 4},
				{1, 3, 4},
				{1, 2, 4},
				{1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)

			cloned := cloneGraph(original)

			if original == nil {
				if cloned != nil {
					t.Fatal("expected nil clone")
				}
				return
			}

			if cloned == nil {
				t.Fatal("expected non-nil clone")
			}

			if !deepEqualGraph(original, cloned) {
				t.Fatalf("graph structure mismatch")
			}

			if !verifyDeepCopy(original, cloned) {
				t.Fatalf("graph is not a deep copy")
			}
		})
	}
}
