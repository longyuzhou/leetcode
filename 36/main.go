package main

import "fmt"

type bitmap struct {
	data int
}

// 初始化，并放入1~n
func newBitmap(n int) *bitmap {
	inst := bitmap{-1 ^ (-1 << n)}
	return &inst
}

func (m *bitmap) exist(n int) bool {
	bits := 1 << (n - 1)
	return m.data&bits == bits
}

func (m *bitmap) remove(n int) {
	if !m.exist(n) {
		return
	}
	m.data = m.data ^ (1 << (n - 1))
}

func isValidSudoku(board [][]byte) bool {
	n := 9
	for i := 0; i < n; i++ {
		// horizontal & vertical
		row := newBitmap(n)
		col := newBitmap(n)
		for j := 0; j < n; j++ {
			ch := board[i][j]
			if ch != '.' {
				num := int(ch - '0')
				if !row.exist(num) {
					return false
				}
				row.remove(num)
			}

			ch = board[j][i]
			if ch != '.' {
				num := int(ch - '0')
				if !col.exist(num) {
					return false
				}
				col.remove(num)
			}
		}

		// sub-boxes
		x, y := 3*(i/3), i%3*3
		subBox := newBitmap(n)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				ch := board[x+j][y+k]
				if ch != '.' {
					num := int(ch - '0')
					if !subBox.exist(num) {
						return false
					}
					subBox.remove(num)
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board) == true)

	board = [][]byte{
		[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board) == false)

	board = [][]byte{
		[]byte{'.', '9', '.', '.', '4', '.', '.', '.', '.'},
		[]byte{'1', '.', '.', '.', '.', '.', '6', '.', '.'},
		[]byte{'.', '.', '3', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		[]byte{'3', '.', '.', '.', '5', '.', '.', '.', '.'},
		[]byte{'.', '.', '7', '.', '.', '4', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '7', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(board) == true)
}
