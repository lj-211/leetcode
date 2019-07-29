package solutions

import (
	"reflect"
)

func kthSmallest(matrix [][]int, k int) int {
	return 0
}

func init() {
	desc := `
Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 ≤ k ≤ n2.
	`
	sol := Solution{
		Title:  "Kth Smallest Element in a Sorted Matrix",
		Desc:   desc,
		Method: reflect.ValueOf(kthSmallest),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[][]int{
		[]int{1, 5, 9},
		[]int{10, 11, 13},
		[]int{12, 13, 15},
	}, 8}
	a.Output = []interface{}{13}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0378"] = sol
}
