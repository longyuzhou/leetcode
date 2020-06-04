package main

import (
	"fmt"
)

func matrixBlockSum(mat [][]int, K int) [][]int {
	m := len(mat)
	n := len(mat[0])

	// hor[y][x] = hor[y][x-1] + mat[y][x]
	hor := make([][]int, m)
	for y := 0; y < m; y++ {
		hor[y] = make([]int, n)
		for x := 0; x < n; x++ {
			if x == 0 {
				hor[y][x] = mat[y][x]
			} else {
				hor[y][x] = hor[y][x-1] + mat[y][x]
			}
		}
	}

	// ver[y][x] = ver[y-1][x] + hor[y][x+K] - hor[y][x-K-1]
	ver := make([][]int, m)
	for y := 0; y < m; y++ {
		ver[y] = make([]int, n)
		for x := 0; x < n; x++ {
			h := hor[y][min(x+K, n-1)]
			if x-K-1 >= 0 {
				h -= hor[y][x-K-1]
			}
			if y == 0 {
				ver[y][x] = h
			} else {
				ver[y][x] = ver[y-1][x] + h
			}
		}
	}

	// res[y][x] = ver[y+K][x] - ver[y-K-1][x]
	res := make([][]int, m)
	for y := 0; y < m; y++ {
		res[y] = make([]int, n)
		for x := 0; x < n; x++ {
			res[y][x] += ver[min(y+K, m-1)][x]
			if y-K-1 >= 0 {
				res[y][x] -= ver[y-K-1][x]
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func test(mat [][]int, K int, expect [][]int) {
	actual := matrixBlockSum(mat, K)
	pass := true
	for i := range expect {
		for j := range expect[i] {
			if expect[i][j] != actual[i][j] {
				pass = false
				break
			}
		}
	}
	if pass {
		fmt.Println("pass")
	} else {
		fmt.Printf("fail! expect %v, got %v\n", expect, actual)
	}
}

func main() {
	test(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9}},
		1,
		[][]int{
			[]int{12, 21, 16},
			[]int{27, 45, 33},
			[]int{24, 39, 28},
		},
	)
	test(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9}},
		2,
		[][]int{
			[]int{45, 45, 45},
			[]int{45, 45, 45},
			[]int{45, 45, 45},
		},
	)
}
