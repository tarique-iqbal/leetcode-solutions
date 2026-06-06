package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			result[pos] = leftSq
			left++
		} else {
			result[pos] = rightSq
			right--
		}

		pos--
	}

	return result
}

func main() {
	testCases := [][]int{
		{-4, -1, 0, 3, 10},
		{-7, -3, 2, 3, 11},
		{-5, -4, -3, -2, -1}, // all negative
		{0, 1, 2, 3, 4},      // all non-negative
		{-2, -1, 0, 1, 2},    // symmetric
		{0},                  // single zero
		{-1},                 // single negative
		{5},                  // single positive
		{},                   // empty slice
	}

	for _, nums := range testCases {
		fmt.Printf("Input:  %v\n", nums)
		fmt.Printf("Output: %v\n", sortedSquares(nums))
		fmt.Println()
	}
}
