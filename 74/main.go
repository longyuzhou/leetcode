package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	get := func(i int) int {
		row := i / n
		col := i % n
		return matrix[row][col]
	}

	i, j := 0, m*n-1
	for i < j {
		mid := int(uint(i+j) >> 1)
		num := get(mid)
		if target > num {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return get(i) == target
}

func main() {
	matrix := [][]int{}
	fmt.Println(searchMatrix(matrix, 3) == false)

	matrix = [][]int{[]int{}}
	fmt.Println(searchMatrix(matrix, 3) == false)

	matrix = [][]int{[]int{1}}
	fmt.Println(searchMatrix(matrix, 1) == true)

	matrix = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 3) == true)

	matrix = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 13) == false)
}
