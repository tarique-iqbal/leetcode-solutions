package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelVals = append(levelVals, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelVals)
	}

	return result
}

func main() {
	// Construct a sample binary tree:
	//        3
	//       / \
	//      9  20
	//         / \
	//        15  7
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := levelOrder(root)
	fmt.Println("Level Order Traversal:", result)
}
