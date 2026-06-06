package main

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "LeetCode example 1",
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "LeetCode example 2",
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single negative",
			input:    []int{-5},
			expected: []int{25},
		},
		{
			name:     "Single positive",
			input:    []int{5},
			expected: []int{25},
		},
		{
			name:     "Single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "All negative",
			input:    []int{-5, -4, -3, -2, -1},
			expected: []int{1, 4, 9, 16, 25},
		},
		{
			name:     "All positive",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 9, 16},
		},
		{
			name:     "Symmetric around zero",
			input:    []int{-2, -1, 0, 1, 2},
			expected: []int{0, 1, 1, 4, 4},
		},
		{
			name:     "Two negatives",
			input:    []int{-2, -1},
			expected: []int{1, 4},
		},
		{
			name:     "Duplicate negatives",
			input:    []int{-2, -2},
			expected: []int{4, 4},
		},
		{
			name:     "Duplicate positives",
			input:    []int{3, 3},
			expected: []int{9, 9},
		},
		{
			name:     "Equal magnitude values",
			input:    []int{-3, 3},
			expected: []int{9, 9},
		},
		{
			name:     "Large mixed example",
			input:    []int{-10, -5, -2, 0, 1, 7},
			expected: []int{0, 1, 4, 25, 49, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf(
					"sortedSquares(%v) = %v; want %v",
					tt.input,
					got,
					tt.expected,
				)
			}
		})
	}
}
