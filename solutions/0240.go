package solutions

import (
	"reflect"
)

func searchMatrix(matrix [][]int, target int) bool {
	rsize := len(matrix)
	if rsize == 0 || len(matrix[0]) == 0 {
		return false
	}
	csize := len(matrix[0])
	i := rsize - 1
	j := csize - 1
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}
	for i >= 0 && j >= 0 {
		val := matrix[i][j]
		if val == target {
			return true
		} else if target < val {
			j--
		} else {
			i++
		}
	}

	return false
}

func init() {
	desc := `
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.
	`
	sol := Solution{
		Title:  "Search a 2D Matrix II",
		Desc:   desc,
		Method: reflect.ValueOf(searchMatrix),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0240"] = sol
}
