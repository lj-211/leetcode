/*
 * 79. Word Search
 * Given a 2D board and a word, find if the word exists in the grid.
 * The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
 *
 * Example:
 * board =
 * [
 *   ['A','B','C','E'],
 *   ['S','F','C','S'],
 *   ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 */
package solutions

import (
	"fmt"
)

func internal_exist(board [][]byte, x int, y int, word string, idx int) bool {
	if idx == len(word) {
		return true
	}
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}
	if board[x][y] != word[idx] {
		return false
	}

	board[x][y] ^= 128
	get := internal_exist(board, x-1, y, word, idx+1) ||
		internal_exist(board, x+1, y, word, idx+1) ||
		internal_exist(board, x, y-1, word, idx+1) ||
		internal_exist(board, x, y+1, word, idx+1)
	board[x][y] ^= 128

	return get
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if internal_exist(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	input := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	has := exist(input, "ABCCED")
	fmt.Println(has)
}
