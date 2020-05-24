package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	i, j, k, x, y := 0, 0, n, 0, 0
	for ; i < n*n; i++ {
		res[y][x] = i + 1
		if j < (k-1)*4-1 {
			if j/(k-1) == 0 {
				x++
			} else if j/(k-1) == 1 {
				y++
			} else if j/(k-1) == 2 {
				x--
			} else {
				y--
			}
			j++
		} else {
			j = 0
			k -= 2
			x++
		}
	}
	return res
}

func main() {
	n := 3
	for _, row := range generateMatrix(n) {
		fmt.Println(row)
	}
}
