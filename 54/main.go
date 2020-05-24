package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	row, col := len(matrix), len(matrix[0])
	res := make([]int, row*col)
	i, x, y := 0, 0, 0
	for j := 0; j < len(res); j++ {
		res[j] = matrix[y][x]

		b, c := col-1, row-1+col-1
		if i < c*2-1 {
			if i%c < b {
				x += 1*(1-i/c) + -1*i/c
			} else {
				y += 1*(1-i/c) + -1*i/c
			}
			i++
		} else {
			row -= 2
			col -= 2
			x++
			i = 0
		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		[]int{1, 2, 3, 4},
	}))
	fmt.Println(spiralOrder([][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}))
}
