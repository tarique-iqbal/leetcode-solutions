package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	l := lowerBound(nums, target)

	if l == n || nums[l] != target {
		return []int{-1, -1}
	}

	r := upperBound(nums, target) - 1

	return []int{l, r}
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3 4]

	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1 -1]

	fmt.Println(searchRange([]int{}, 0)) // [-1 -1]

	fmt.Println(searchRange([]int{1}, 1)) // [0 0]

	fmt.Println(searchRange([]int{2, 2}, 2)) // [0 1]

	fmt.Println(searchRange([]int{1, 2, 3}, 2)) // [1 1]
}
