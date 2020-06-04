package main

import (
	"fmt"
)

func kConcatenationMaxSum(arr []int, k int) int {
	size := len(arr)
	if size == 0 {
		return 0
	}

	modulo := 1000000007

	// when k == 1
	sum := 0
	maxSum := 0

	// from left to right
	leftSum := 0
	maxLeftSum := 0

	// from right to left
	rightSum := 0
	maxRightSum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
		sum = max(sum, 0)
		maxSum = max(maxSum, sum)

		leftSum += arr[i]
		maxLeftSum = max(maxLeftSum, leftSum)

		rightSum += arr[size-1-i]
		maxRightSum = max(maxRightSum, rightSum)
	}

	res := max(maxSum, leftSum*k)
	if k >= 2 {
		if leftSum <= 0 {
			res = max(res, maxRightSum+maxLeftSum)
		} else {
			res = max(res, maxRightSum+leftSum*(k-2)+maxLeftSum)
		}
	}
	return res % modulo
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(kConcatenationMaxSum([]int{1, 2}, 3) == 9)
	fmt.Println(kConcatenationMaxSum([]int{1, -2, 1}, 5) == 2)
	fmt.Println(kConcatenationMaxSum([]int{-1, -2}, 7) == 0)
	fmt.Println(kConcatenationMaxSum([]int{-5, -2, 0, 0, 3, 9, -2, -5, 4}, 5) == 20)
}
