package main

import "fmt"

// Given a triangle, find the minimum path sum from top to bottom.
// Each step you may move to adjacent numbers on the row below.
// For example, given the following triangle
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
// Note:
// Bonus point if you are able to do this using only O(n) extra space,
// where n is the total number of rows in the triangle.
//
// Solution: 最短路径问题, 使用动态规划求解
//   dp[n] = min(dp[n-1], dp[n]) + v(n)
//   注意 n-1, n 越界情况
//   因为只使用O(n)空间，上一层的dp[n]会被当前层的dp[n]覆盖，所需须将上一层的dp[n]存到临时变量
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for i, row := range triangle {
		upLeft := 0
		for j, val := range row {
			tmp := dp[j]
			if j == 0 {
				dp[j] += val
			} else if j == i {
				dp[j] = upLeft + val
			} else {
				dp[j] = min(upLeft, dp[j]) + val
			}
			upLeft = tmp
		}
	}
	return min(dp...)
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("illegal argument")
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res
}

func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle) == 11)
}
