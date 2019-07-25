package solutions

import (
	"reflect"
)

func Compress(old int, new int) int {
	return new*10 + old
}

func DepressNew(new int) int {
	return new / 10
}

func DepressOld(old int) int {
	return old % 10
}

func checkElement(board [][]int, m, n int, i, j int, isOld bool) {
	// i-1, j-1 -> i+1, j+1
	si := i
	sj := j
	ei := i
	ej := j
	if i-1 >= 0 {
		si = i - 1
	}
	if i+1 < m {
		ei = i + 1
	}
	if j-1 >= 0 {
		sj = j - 1
	}
	if j+1 < n {
		ej = j + 1
	}

	cnt := 0
	cur := DepressOld(board[i][j])
	for ii := si; ii <= ei; ii++ {
		for jj := sj; jj <= ej; jj++ {
			if ii == i && jj == j {
				continue
			}
			old := board[ii][jj]
			old = DepressOld(old)
			if old == 1 {
				cnt++
			}
		}
	}

	new := 0
	if cur == 1 && cnt < 2 {
		new = 0
	} else if cur == 1 && (cnt == 2 || cnt == 3) {
		new = 1
	} else if cur == 1 && cnt > 3 {
		new = 0
	} else if cur == 0 && cnt == 3 {
		new = 1
	}

	board[i][j] = Compress(cur, new)
}

func gameOfLife(board [][]int) {
	m := len(board)
	n := 0
	if m > 0 {
		n = len(board[0])
	}
	if m == 0 && n == 0 {
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			checkElement(board, m, n, i, j, true)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = DepressNew(board[i][j])
		}
	}

}

func init() {
	desc := `
According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population..
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
Write a function to compute the next state (after one update) of the board given its current state. The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously.

Example:

Input:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
Output:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
Follow up:

Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells first and then use their updated values to update other cells.
In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array. How would you address these problems?
	`
	sol := Solution{
		Title:  "Game of Life",
		Desc:   desc,
		Method: reflect.ValueOf(gameOfLife),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}}
	a.Output = []interface{}{}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0289"] = sol
}
