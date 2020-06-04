package main

import "fmt"

// Given a list of non-negative numbers and a target integer k,
// write a function to check if the array
// has a continuous subarray of size at least 2
// that sums up to a multiple of k,
// that is, sums up to n*k where n is also an integer.
//
// Example 1:
//   Input: [23, 2, 4, 6, 7],  k=6
//   Output: True
//   Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
//
// Solution:
//       [23,  2,  4,  6,  7]
//   sum [23, 25, 29, 35, 42]
//   % 6 [ 5,  1,  5,  5,  0] // dp
//     dp[i] == dp[j] && j-i >= 2
//  or dp[i] == 0
//
// 注意k == 0的情况
func checkSubarraySum(nums []int, k int) bool {
	minSize := 2
	dp := make(map[int]int, len(nums))
	total := 0
	for i, num := range nums {
		total += num
		if ((k == 0 && total == 0) || (k != 0 && total%k == 0)) && i+1 >= minSize {
			return true
		}
		key := total
		if k != 0 {
			key = total % k
		}
		if idx, ok := dp[key]; ok && i-idx >= minSize {
			return true
		} else if !ok {
			dp[key] = i
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{0, 0}, 0) == true)
	fmt.Println(checkSubarraySum([]int{0, 1, 0}, 0) == false)
	fmt.Println(checkSubarraySum([]int{-1, 1}, 0) == true)
	fmt.Println(checkSubarraySum([]int{1, -2, 1}, 0) == true)
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 0) == false)
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6) == true)
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6) == true)
}
