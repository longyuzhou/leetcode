package main

import "fmt"

// Given a 2D board containing 'X' and 'O' (the letter O),
// capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
// Example:
//   X X X X
//   X O O X
//   X X O X
//   X O X X
// After running your function, the board should be:
//   X X X X
//   X X X X
//   X X X X
//   X O X X

func solve(board [][]byte) {
	h := len(board)
	if h <= 1 {
		return
	}
	w := len(board[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		board[y][x] = '*'
		if y > 0 && board[y-1][x] == 'O' {
			dfs(x, y-1)
		}
		if x < w-1 && board[y][x+1] == 'O' {
			dfs(x+1, y)
		}
		if y < h-1 && board[y+1][x] == 'O' {
			dfs(x, y+1)
		}
		if x > 0 && board[y][x-1] == 'O' {
			dfs(x-1, y)
		}
	}

	for y, row := range board {
		for x, val := range row {
			if val == 'O' && (y == 0 || y == h-1 || x == 0 || x == w-1) {
				dfs(x, y)
			}
		}
	}

	for y, row := range board {
		for x, val := range row {
			if val == '*' {
				board[y][x] = 'O'
			} else if val == 'O' {
				board[y][x] = 'X'
			}
		}
	}
}

func test(board [][]byte, expect [][]byte) {
	solve(board)
	pass := true
	for i := 0; i < len(expect); i++ {
		for j := 0; j < len(expect[i]); j++ {
			if expect[i][j] != board[i][j] {
				pass = false
				break
			}
		}
	}
	if !pass {
		fmt.Printf("fail! expect %v, got %v\n", expect, board)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test(
		[][]byte{
			[]byte{'O', 'O'},
			[]byte{'O', 'O'},
		},
		[][]byte{
			[]byte{'O', 'O'},
			[]byte{'O', 'O'},
		},
	)

	test(
		[][]byte{
			[]byte{'O', 'O', 'O'},
			[]byte{'O', 'O', 'O'},
			[]byte{'O', 'O', 'O'},
		},
		[][]byte{
			[]byte{'O', 'O', 'O'},
			[]byte{'O', 'O', 'O'},
			[]byte{'O', 'O', 'O'},
		},
	)

	test(
		[][]byte{
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'O', 'O', 'X'},
			[]byte{'X', 'X', 'O', 'X'},
			[]byte{'X', 'O', 'X', 'X'},
		},
		[][]byte{
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'O', 'X', 'X'},
		},
	)

	solve([][]byte{
		[]byte{'O', 'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'O'},
		[]byte{'X', 'O', 'X', 'O', 'X'},
		[]byte{'O', 'X', 'O', 'O', 'O'},
		[]byte{'X', 'X', 'O', 'X', 'O'},
	})
}
