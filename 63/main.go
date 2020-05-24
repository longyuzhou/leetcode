package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
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
			if obstacleGrid[i][j] != 0 {
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = get(i-1, j) + get(i, j-1)
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid) == 2)
}
