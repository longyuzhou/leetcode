package main

import (
	"fmt"
)

// Given a non-empty array of integers,
// every element appears three times except for one,
// which appears exactly once. Find that single one.
//
// Note:
// Your algorithm should have a linear runtime complexity.
// Could you implement it without using extra memory?
// Example 1:
//   Input: [2,2,3,2]
//   Output: 3
// Example 2:
//   Input: [0,1,0,1,0,1,99]
//   Output: 99
//
// Solution:
//   2 = 0 0 1 0
//   2 = 0 0 1 0
//   4 = 0 1 0 0
//   2 = 0 0 1 0
//       0 1 3 0 统计1出现的次数，然后%3
//  %3 = 0 1 0 0 即问题的解为4
// 注意精度溢出
func singleNumber(nums []int) int {
	counts := make([]int, 32)
	for _, num := range nums {
		for i := range counts {
			mask := 1 << i
			if num&mask == mask {
				counts[i]++
			}
		}
	}

	res := int32(0)
	for i, count := range counts {
		if count%3 != 0 {
			res = res | 1<<i
		}
	}
	return int(res)
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}) == 3)
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}) == 99)
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}) == -4)
}
