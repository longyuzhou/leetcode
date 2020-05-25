package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	var dfs func(x, y, start int, path [][]int) bool
	dfs = func(x, y, start int, path [][]int) bool {
		if start >= len(word) {
			return true
		}

		// 检查坐标边界，以及当前坐标处元素是否与要搜索的元素相同
		if y >= len(board) || y < 0 ||
			x >= len(board[0]) || x < 0 ||
			start >= len(word) ||
			board[y][x] != word[start] {
			return false
		}

		// 同一个坐标不能使用多次
		for _, pos := range path {
			if pos[0] == x && pos[1] == y {
				return false
			}
		}

		path = append(path, []int{x, y})

		// 递归搜索左、右、上、下
		return dfs(x-1, y, start+1, path) ||
			dfs(x+1, y, start+1, path) ||
			dfs(x, y-1, start+1, path) ||
			dfs(x, y+1, start+1, path)
	}

	// 先找到word[0]在矩阵中的位置
	for i, row := range board {
		for j := range row {
			if board[i][j] == word[0] {
				if dfs(j, i, 0, [][]int{}) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	fmt.Println(exist([][]byte{}, "A") == false)
	fmt.Println(exist([][]byte{[]byte{'A', 'B', 'C', 'E'}}, "A") == true)
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED") == true)
	fmt.Println(exist(board, "SEE") == true)
	fmt.Println(exist(board, "ABCB") == false)
}
