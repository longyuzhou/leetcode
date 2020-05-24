package main

import "fmt"

/**
 * first, do a diagonally fold
 *   1, 2, 3      9, 6, 3
 *   4, 5, 6  =>  8, 5, 2
 *   7, 8, 9      7, 4, 1
 * then, do a vertically fold
 *   9, 6, 3      7, 4, 1
 *   8, 5, 2  =>  8, 5, 2
 *   7, 4, 1      9, 6, 3
 */
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	if n <= 0 {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-j][n-i] = matrix[n-j][n-i], matrix[i][j]
		}
	}

	for i := 0; i <= n; i++ {
		j, k := 0, n
		for j < k {
			matrix[j][i], matrix[k][i] = matrix[k][i], matrix[j][i]
			j++
			k--
		}
	}
}

func test(input, output [][]int) {
	rotate(input)
	pass := true
	for i, r1 := range input {
		r2 := output[i]
		for j := range r1 {
			if r1[j] != r2[j] {
				pass = false
				break
			}
		}
	}
	if !pass {
		fmt.Println("fail!", input)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}, [][]int{
		[]int{7, 4, 1},
		[]int{8, 5, 2},
		[]int{9, 6, 3},
	})

	test([][]int{[]int{0}}, [][]int{[]int{0}})

	test([][]int{}, [][]int{})
}
