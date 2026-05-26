package main

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name     string
		times    [][]int
		n        int
		k        int
		expected int
	}{
		{
			name: "basic example",
			times: [][]int{
				{2, 1, 1},
				{2, 3, 1},
				{3, 4, 1},
			},
			n:        4,
			k:        2,
			expected: 2,
		},
		{
			name:     "single node",
			times:    [][]int{},
			n:        1,
			k:        1,
			expected: 0,
		},
		{
			name: "unreachable node",
			times: [][]int{
				{1, 2, 1},
			},
			n:        3,
			k:        1,
			expected: -1,
		},
		{
			name: "linear graph",
			times: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{3, 4, 3},
			},
			n:        4,
			k:        1,
			expected: 6,
		},
		{
			name: "multiple paths choose shortest",
			times: [][]int{
				{1, 2, 1},
				{1, 3, 4},
				{2, 3, 2},
			},
			n:        3,
			k:        1,
			expected: 3,
		},
		{
			name: "cycle graph",
			times: [][]int{
				{1, 2, 1},
				{2, 3, 1},
				{3, 1, 1},
			},
			n:        3,
			k:        1,
			expected: 2,
		},
		{
			name: "disconnected graph",
			times: [][]int{
				{1, 2, 1},
				{3, 4, 1},
			},
			n:        4,
			k:        1,
			expected: -1,
		},
		{
			name: "direct path better than indirect",
			times: [][]int{
				{1, 2, 10},
				{1, 3, 1},
				{3, 2, 1},
			},
			n:        3,
			k:        1,
			expected: 2,
		},
		{
			name: "all nodes directly connected",
			times: [][]int{
				{1, 2, 1},
				{1, 3, 2},
				{1, 4, 3},
			},
			n:        4,
			k:        1,
			expected: 3,
		},
		{
			name: "duplicate edges choose minimum",
			times: [][]int{
				{1, 2, 5},
				{1, 2, 2},
				{2, 3, 1},
			},
			n:        3,
			k:        1,
			expected: 3,
		},
		{
			name: "large edge weights",
			times: [][]int{
				{1, 2, 100},
				{2, 3, 200},
			},
			n:        3,
			k:        1,
			expected: 300,
		},
		{
			name: "starting node isolated",
			times: [][]int{
				{2, 3, 1},
			},
			n:        3,
			k:        1,
			expected: -1,
		},
		{
			name: "fully connected graph",
			times: [][]int{
				{1, 2, 1},
				{1, 3, 5},
				{2, 3, 1},
				{3, 4, 1},
				{2, 4, 4},
			},
			n:        4,
			k:        1,
			expected: 3,
		},
		{
			name:     "zero edges multiple nodes",
			times:    [][]int{},
			n:        5,
			k:        1,
			expected: -1,
		},
		{
			name: "single edge reachable",
			times: [][]int{
				{1, 2, 7},
			},
			n:        2,
			k:        1,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := networkDelayTime(tt.times, tt.n, tt.k)

			if result != tt.expected {
				t.Errorf(
					"networkDelayTime(%v, %d, %d) = %d; want %d",
					tt.times,
					tt.n,
					tt.k,
					result,
					tt.expected,
				)
			}
		})
	}
}
