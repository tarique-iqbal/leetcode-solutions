package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "example case",
			root: &TreeNode{3,
				&TreeNode{9, nil, nil},
				&TreeNode{20,
					&TreeNode{15, nil, nil},
					&TreeNode{7, nil, nil},
				},
			},
			expected: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name:     "empty tree",
			root:     nil,
			expected: nil,
		},
		{
			name: "single node",
			root: &TreeNode{1, nil, nil},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "tree with two levels",
			root: &TreeNode{1,
				&TreeNode{2, nil, nil},
				&TreeNode{3, nil, nil},
			},
			expected: [][]int{
				{1},
				{2, 3},
			},
		},
		{
			name: "tree with more than two levels",
			root: &TreeNode{1,
				&TreeNode{2,
					&TreeNode{4, nil, nil},
					&TreeNode{5, nil, nil},
				},
				&TreeNode{3,
					nil,
					&TreeNode{6, nil, nil},
				},
			},
			expected: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.expected)
			}
		})
	}
}
