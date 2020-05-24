package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	get := func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 0
		}
		return dp[x][y]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			up := get(i-1, j)
			left := get(i, j-1)
			dp[i][j] = grid[i][j]
			if i == 0 && j == 0 {
				// ignore
			} else if i == 0 {
				dp[i][j] += left
			} else if j == 0 {
				dp[i][j] += up
			} else {
				dp[i][j] += min(up, left)
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	fmt.Println(minPathSum(grid) == 7)
}
