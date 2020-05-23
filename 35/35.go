package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	return search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}

func search(n int, f func(n int) bool) int {
	i, j := 0, n
	for i < j {
		mid := int(uint(i+j) >> 1)
		if !f(mid) {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func test(nums []int, target, expect int) {
	fmt.Println("Input:", nums, target)
	fmt.Println("Expect:", expect)
	actual := searchInsert(nums, target)
	fmt.Println("Actutal:", actual)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
	fmt.Println("====================")
}

func main() {
	test([]int{1, 3, 5, 6}, 5, 2)
	test([]int{1, 3, 5, 6}, 2, 1)
	test([]int{1, 3, 5, 6}, 7, 4)
	test([]int{1, 3, 5, 6}, 0, 0)
}
