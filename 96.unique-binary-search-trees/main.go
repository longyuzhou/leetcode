package main

import (
	"fmt"
)

// 同95.unique-binary-search-trees-ii, 本题使用一维动态规划降低时间复杂度
func numTrees(n int) int {
	if n < 1 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		count := 0
		for j := 0; j < i; j++ {
			count += dp[j] * dp[i-1-j]
		}
		dp[i] = count
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(0) == 0)
	fmt.Println(numTrees(1) == 1)
	fmt.Println(numTrees(2) == 2)
	fmt.Println(numTrees(3) == 5)
}
