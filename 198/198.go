package main

import "fmt"

// 一维动态规划
func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return nums[0]
	}

	dp := make([]int, size+1)
	dp[1] = nums[0]
	for i := 2; i <= size; i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	return dp[size]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func test(nums []int, expect int) {
	actual := rob(nums)
	if actual != expect {
		fmt.Printf("fail! input: %d expect: %d actual: %d\n", nums, expect, actual)
	}
}

func main() {
	test([]int{}, 0)
	test([]int{1}, 1)
	test([]int{2, 3, 2}, 4)
	test([]int{1, 2, 3, 1}, 4)
	test([]int{2, 7, 9, 3, 1}, 12)
}
