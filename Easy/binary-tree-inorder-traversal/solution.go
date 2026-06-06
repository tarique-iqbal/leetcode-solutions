package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	curr := root

	for curr != nil || len(stack) > 0 {
		// go all the way left
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// pop
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// visit root
		result = append(result, curr.Val)

		// traverse right subtree
		curr = curr.Right
	}

	return result
}

func main() {
	// Example 1:
	//
	//     1
	//      \
	//       2
	//      /
	//     3
	//
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println("Example 1:", inorderTraversal(root1))
	// Output: [1 3 2]

	// Example 2:
	//
	//         4
	//       /   \
	//      2     6
	//     / \   / \
	//    1   3 5   7
	//
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println("Example 2:", inorderTraversal(root2))
	// Output: [1 2 3 4 5 6 7]

	// Example 3: Single node
	root3 := &TreeNode{
		Val: 42,
	}

	fmt.Println("Example 3:", inorderTraversal(root3))
	// Output: [42]

	// Example 4: Empty tree
	var root4 *TreeNode

	fmt.Println("Example 4:", inorderTraversal(root4))
	// Output: []
}
