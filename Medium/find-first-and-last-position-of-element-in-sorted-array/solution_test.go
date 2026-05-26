package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "empty array",
			nums:   []int{},
			target: 5,
			want:   []int{-1, -1},
		},
		{
			name:   "single element - found",
			nums:   []int{5},
			target: 5,
			want:   []int{0, 0},
		},
		{
			name:   "single element - not found",
			nums:   []int{5},
			target: 3,
			want:   []int{-1, -1},
		},
		{
			name:   "multiple occurrences in middle",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "target not present",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "all elements equal - target present",
			nums:   []int{2, 2, 2, 2, 2},
			target: 2,
			want:   []int{0, 4},
		},
		{
			name:   "all elements equal - target absent",
			nums:   []int{2, 2, 2, 2},
			target: 3,
			want:   []int{-1, -1},
		},
		{
			name:   "target at start",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1,
			want:   []int{0, 0},
		},
		{
			name:   "target at end",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			want:   []int{4, 4},
		},
		{
			name:   "larger array with two blocks",
			nums:   []int{1, 1, 2, 3, 3, 3, 4, 5, 5},
			target: 3,
			want:   []int{3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
