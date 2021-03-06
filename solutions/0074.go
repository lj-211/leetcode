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

	s := 0
	e := rsize*csize - 1
	for s <= e {
		mid := (s + e) / 2
		mval := matrix[mid/csize][mid%csize]
		if mval == target {
			return true
		} else if mval > target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}

	return false
}

func init() {
	desc := `
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
	`
	sol := Solution{
		Title:  "Search a 2D Matrix",
		Desc:   desc,
		Method: reflect.ValueOf(searchMatrix),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0074"] = sol
}
