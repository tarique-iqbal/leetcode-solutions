package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	queue := []*Node{node}
	visited[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, nei := range cur.Neighbors {
			if _, ok := visited[nei]; !ok {
				visited[nei] = &Node{Val: nei.Val}
				queue = append(queue, nei)
			}

			visited[cur].Neighbors = append(
				visited[cur].Neighbors,
				visited[nei],
			)
		}
	}

	return visited[node]
}

func printGraph(node *Node) {
	if node == nil {
		fmt.Println("empty graph")
		return
	}

	visited := make(map[*Node]bool)
	queue := []*Node{node}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visited[cur] {
			continue
		}

		visited[cur] = true

		fmt.Printf("Node %d -> [", cur.Val)

		for i, nei := range cur.Neighbors {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(nei.Val)

			if !visited[nei] {
				queue = append(queue, nei)
			}
		}

		fmt.Println("]")
	}
}

func main() {
	// Build graph:
	// 1 -- 2
	// |    |
	// 4 -- 3

	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	fmt.Println("Original Graph:")
	printGraph(n1)

	clone := cloneGraph(n1)

	fmt.Println("\nCloned Graph:")
	printGraph(clone)

	fmt.Println("\nDeep Copy Verification:")
	fmt.Printf("Original Node 1 Address: %p\n", n1)
	fmt.Printf("Cloned   Node 1 Address: %p\n", clone)

	fmt.Printf("Same pointer? %v\n", n1 == clone)
}
