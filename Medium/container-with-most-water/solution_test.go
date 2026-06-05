package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "LeetCode example",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Minimum valid input",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Symmetric walls",
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			name:     "Small valley",
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			name:     "Mixed heights",
			height:   []int{2, 3, 10, 5, 7, 8, 9},
			expected: 36,
		},
		{
			name:     "Empty slice",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			height:   []int{5},
			expected: 0,
		},
		{
			name:     "Two different heights",
			height:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "All heights equal",
			height:   []int{5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "Strictly increasing",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Strictly decreasing",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6,
		},
		{
			name:     "Outer walls produce max area",
			height:   []int{10, 1, 1, 1, 10},
			expected: 40,
		},
		{
			name:     "Tallest walls not widest",
			height:   []int{1, 8, 100, 2, 100, 8, 1},
			expected: 200,
		},
		{
			name:     "Multiple equal maximums",
			height:   []int{3, 3, 3, 3},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.height)

			if result != tt.expected {
				t.Fatalf(
					"maxArea(%v) = %d; want %d",
					tt.height,
					result,
					tt.expected,
				)
			}
		})
	}
}
