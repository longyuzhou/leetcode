package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	xor := 0
	for i, num := range nums {
		xor = xor ^ (i + 1) ^ num
	}
	return xor
}

func test(nums []int, expect int) {
	actual := missingNumber(nums)
	if actual != expect {
		fmt.Printf("fail! Input: %v, Expect: %d\n", nums, expect)
	}
}

func main() {
	test([]int{3, 0, 1}, 2)
	test([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8)
}
