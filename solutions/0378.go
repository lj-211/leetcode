package solutions

import (
	"fmt"
	"reflect"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	min := matrix[0][0]
	max := matrix[n-1][n-1]

	for min <= max {
		cnt := 0
		j := n - 1
		mid := (min + max) / 2
		fmt.Println("搜索区间: ", min, max, mid)
		maxVal := min
		for i := 0; i < n; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			cnt += (j + 1)
			if j >= 0 && matrix[i][j] > maxVal {
				maxVal = matrix[i][j]
			}
		}

		fmt.Println("cnt: ", cnt)

		if cnt == k {
			return maxVal
		} else if cnt < k {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return min
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
	/*
		a.Input = []interface{}{[][]int{
			[]int{1, 5, 9},
			[]int{10, 11, 13},
			[]int{12, 13, 15},
		}, 8}
		a.Output = []interface{}{13}
		sol.Tests = append(sol.Tests, a)
	*/
	a.Input = []interface{}{[][]int{
		[]int{1, 2},
		[]int{3, 3},
	}, 2}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0378"] = sol
}
