package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxCurrent, maxGlobal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxCurrent = max(nums[i], maxCurrent+nums[i])
		maxGlobal = max(maxGlobal, maxCurrent)
	}
	return maxGlobal
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func test(nums []int, expect int) {
	actual := maxSubArray(nums)
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Expect: %d\n", expect)
	fmt.Printf("Output: %d\n", actual)
	if actual != expect {
		fmt.Println("fail")
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test([]int{-2, 1}, 1)
	test([]int{-2, -1}, -1)
	test([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6)
}
