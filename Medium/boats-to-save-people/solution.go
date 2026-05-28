package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	left, right := 0, len(people)-1
	boats := 0

	for left <= right {
		// lightest + heaviest can share a boat
		if people[left]+people[right] <= limit {
			left++
			right--
		} else {
			// heaviest goes alone
			right--
		}

		boats++
	}

	return boats
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3)) // output: 1

	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3)) // output: 3

	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5)) // output: 4

	fmt.Println(numRescueBoats([]int{1, 1, 2, 2}, 3)) // output: 2
}
