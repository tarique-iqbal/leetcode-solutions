package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Nil tree",
			root:     nil,
			expected: nil,
		},
		{
			name: "Single node",
			root: &TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
		{
			name: "Root with left child",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			expected: []int{1, 2},
		},
		{
			name: "Root with right child",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{1, 2},
		},
		{
			name: "LeetCode example",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name: "Balanced tree",
			root: &TreeNode{
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
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Left skewed tree",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "Right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "Complete tree",
			root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 12,
					Left: &TreeNode{
						Val: 10,
					},
					Right: &TreeNode{
						Val: 14,
					},
				},
			},
			expected: []int{2, 4, 6, 8, 10, 12, 14},
		},
		{
			name: "Duplicate values",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{2, 2, 2},
		},
		{
			name: "Negative values",
			root: &TreeNode{
				Val: -2,
				Left: &TreeNode{
					Val: -3,
				},
				Right: &TreeNode{
					Val: -1,
				},
			},
			expected: []int{-3, -2, -1},
		},
		{
			name: "Unbalanced tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: []int{3, 4, 5, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversal(tt.root)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf(
					"inorderTraversal() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}
