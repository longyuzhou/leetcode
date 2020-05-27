package main

import (
	"fmt"
	"strconv"
)

// A message containing letters from A-Z is being encoded
// to numbers using the following mapping:
//   'A' -> 1
//   'B' -> 2
//   ...
//   'Z' -> 26
// Given a non-empty string containing only digits,
// determine the total number of ways to decode it.
// Example 1:
//   Input: "12"
//   Output: 2
//   Explanation: It could be decoded as "AB" (1 2) or "L" (12).
//
// Solution: 一维动态规划(注意临界情况)
func numDecodings(s string) int {
	size := len(s)
	if size == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, size+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < size; i++ {
		// 当s[i-1:i+1]为"00", "30", "40", "50" ... 时，无解
		if (s[i-1] == '0' || s[i-1] > '2') && s[i] == '0' {
			return 0
		}

		// 排除s[i-1:i+1]为"01" ~ "09", "10", "20" 的情况
		if s[i-1] != '0' && s[i] != '0' {
			dp[i+1] += dp[i]
		}

		// 排除s[i-1:i+1]为"00", "27", "28", "29" ... 的情况
		if num, _ := strconv.Atoi(s[i-1 : i+1]); 0 < num && num <= 26 {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[size]
}

func main() {
	fmt.Println(numDecodings("") == 0)
	fmt.Println(numDecodings("0") == 0)
	fmt.Println(numDecodings("1") == 1)
	fmt.Println(numDecodings("12") == 2)
	fmt.Println(numDecodings("226") == 3)
	fmt.Println(numDecodings("227") == 2)
	fmt.Println(numDecodings("230") == 0)
	fmt.Println(numDecodings("100") == 0)
	fmt.Println(numDecodings("101") == 1)
	fmt.Println(numDecodings("301") == 0)
}
