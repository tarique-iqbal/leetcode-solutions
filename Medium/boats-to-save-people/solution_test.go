package main

import "testing"

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name     string
		people   []int
		limit    int
		expected int
	}{
		{
			name:     "Example case 1",
			people:   []int{1, 2},
			limit:    3,
			expected: 1,
		},
		{
			name:     "Example case 2",
			people:   []int{3, 2, 2, 1},
			limit:    3,
			expected: 3,
		},
		{
			name:     "Example case 3",
			people:   []int{3, 5, 3, 4},
			limit:    5,
			expected: 4,
		},
		{
			name:     "Empty list",
			people:   []int{},
			limit:    10,
			expected: 0,
		},
		{
			name:     "Single person",
			people:   []int{2},
			limit:    3,
			expected: 1,
		},
		{
			name:     "All pairs fit exactly",
			people:   []int{2, 2, 2, 2},
			limit:    4,
			expected: 2,
		},
		{
			name:     "All people require separate boats",
			people:   []int{5, 5, 5, 5},
			limit:    5,
			expected: 4,
		},
		{
			name:     "Odd number of people",
			people:   []int{1, 2, 2},
			limit:    3,
			expected: 2,
		},
		{
			name:     "Reverse sorted input",
			people:   []int{5, 4, 3, 2, 1},
			limit:    5,
			expected: 3,
		},
		{
			name:     "Already sorted input",
			people:   []int{1, 2, 3, 4, 5},
			limit:    5,
			expected: 3,
		},
		{
			name:     "Pair only smallest with largest",
			people:   []int{1, 1, 4, 4},
			limit:    5,
			expected: 2,
		},
		{
			name:     "No pair can fit",
			people:   []int{2, 3, 4},
			limit:    4,
			expected: 3,
		},
		{
			name:     "All under half limit",
			people:   []int{1, 1, 1, 1},
			limit:    10,
			expected: 2,
		},
		{
			name:     "Duplicate weights",
			people:   []int{3, 3, 3, 3, 3},
			limit:    6,
			expected: 3,
		},
		{
			name:     "Exact limit pairing",
			people:   []int{1, 4, 2, 3},
			limit:    5,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy slice to avoid mutation side-effects
			people := append([]int(nil), tt.people...)

			result := numRescueBoats(people, tt.limit)

			if result != tt.expected {
				t.Errorf(
					"numRescueBoats(%v, %d) = %d; want %d",
					tt.people,
					tt.limit,
					result,
					tt.expected,
				)
			}
		})
	}
}
