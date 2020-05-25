package main

import "fmt"

func setZeroes(matrix [][]int) {
	firstRowHasZero := false
	for col, num := range matrix[0] {
		if num == 0 {
			firstRowHasZero = true // 记录第一行是否有0，暂不处理
		}
		// 如果该列有0，则将第一行的该列置为0，作为该列是否应该置为0的标志位
		currColHasZero := false
		for row := 1; row < len(matrix); row++ {
			if matrix[row][col] == 0 {
				currColHasZero = true
				break
			}
		}
		if currColHasZero {
			matrix[0][col] = 0
		}
	}

	for i := 1; i < len(matrix); i++ {
		// 当前行是否有0
		currRowHasZero := false
		for _, num := range matrix[i] {
			if num == 0 {
				currRowHasZero = true
				break
			}
		}
		// 如果当前行有0，或者第1行的相同列为0，则置为0
		for j := range matrix[i] {
			if currRowHasZero || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 如果第一行有0，则全置为0
	if firstRowHasZero {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
}

func print(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	setZeroes(matrix)
	print(matrix)

	matrix = [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	setZeroes(matrix)
	print(matrix)

	matrix = [][]int{
		[]int{1},
		[]int{3},
		[]int{0},
	}
	setZeroes(matrix)
	print(matrix)
}
