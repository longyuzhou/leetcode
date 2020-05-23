package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]^nums[i+1] == 0 {
			return true
		}
	}
	return false
}

func test(nums []int, expect bool) {
	actual := containsDuplicate(nums)
	if actual != expect {
		fmt.Printf("fail! input: %d expect: %t\n", nums, expect)
	}
}

func main() {
	test([]int{}, false)
	test([]int{1}, false)
	test([]int{1, 2, 3, 1}, true)
	test([]int{1, 2, 3, 4}, false)
}
