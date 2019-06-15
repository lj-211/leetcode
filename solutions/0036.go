package solutions

import (
	"fmt"
	"reflect"
)

func isValidSudoku(board [][]byte) bool {
	kmap := make(map[string]bool)

	for k, v := range board {
		for k2, v2 := range v {
			if v2 == '.' {
				continue
			}
			rval := fmt.Sprintf("row-%d-%d", k, v2)
			cval := fmt.Sprintf("col-%d-%d", k2, v2)
			sval := fmt.Sprintf("s-%d-%d-%d", k/3, k2/3, v2)
			if _, ok := kmap[rval]; ok {
				return false
			} else {
				kmap[rval] = true
			}
			if _, ok := kmap[cval]; ok {
				return false
			} else {
				kmap[cval] = true
			}
			if _, ok := kmap[sval]; ok {
				return false
			} else {
				kmap[sval] = true
			}
		}
	}
	return true
}

func init() {
	desc := `
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
A partially filled sudoku which is valid.
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.
	`
	sol := Solution{
		Title:  "Valid Sudoku",
		Desc:   desc,
		Method: reflect.ValueOf(isValidSudoku),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0036"] = sol
}
