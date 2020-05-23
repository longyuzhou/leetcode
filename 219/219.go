package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= k && i+j < len(nums); j++ {
			if nums[i]^nums[i+j] == 0 {
				return true
			}
		}
	}
	return false
}

func test(nums []int, k int, expect bool) {
	actual := containsNearbyDuplicate(nums, k)
	if actual != expect {
		fmt.Printf("fail! input: %d expect: %t\n", nums, expect)
	}
}

func main() {
	test([]int{}, 3, false)
	test([]int{1}, 3, false)
	test([]int{1, 2, 3, 1}, 3, true)
	test([]int{1, 0, 1, 1}, 1, true)
	test([]int{1, 2, 3, 1, 2, 3}, 2, false)
}
