package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left

		minHeight := 0
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}

		area := width * minHeight

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
		{
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			height:   []int{2, 3, 10, 5, 7, 8, 9},
			expected: 36,
		},
	}

	for i, tc := range tests {
		result := maxArea(tc.height)

		fmt.Printf(
			"Test %d: height=%v => result=%d, expected=%d, pass=%v\n",
			i+1,
			tc.height,
			result,
			tc.expected,
			result == tc.expected,
		)
	}
}
