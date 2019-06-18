package solutions

import (
	"reflect"
)

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		val := board[3*(row/3)+i/3][3*(col/3)+i%3]
		if val != '.' && val == c {
			return false
		}
	}

	return true
}

func solve(board [][]byte) bool {
	size := len(board)
	if size == 0 {
		return false
	}
	jsize := len(board[0])
	for i := 0; i < size; i++ {
		for j := 0; j < jsize; j++ {
			if board[i][j] != '.' {
				continue
			}
			for c := byte('1'); c <= byte('9'); c++ {
				if isValid(board, i, j, byte(c)) {
					board[i][j] = c
					if solve(board) {
						return true
					} else {
						board[i][j] = '.'
					}
				}
			}

			return false
		}
	}

	return true
}

func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}
	solve(board)
}

func init() {
	desc := `
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.
Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.
	`
	sol := Solution{
		Title:  "Sudoku Solver",
		Desc:   desc,
		Method: reflect.ValueOf(solveSudoku),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0037"] = sol
}
