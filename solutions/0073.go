package solutions

import (
	"reflect"
)

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	if (m * n) == 0 {
		return
	}
	col0 := 1

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
}

func init() {
	desc := `
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:

Input:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
Example 2:

Input:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
	`
	sol := Solution{
		Title:  "Set Matrix Zeroes",
		Desc:   desc,
		Method: reflect.ValueOf(setZeroes),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}}
	a.Output = []interface{}{[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 4, 5, 0},
		[]int{0, 3, 1, 0},
	}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0073"] = sol
}
