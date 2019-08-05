package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = ds.MinInt(ds.MinInt(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

func init() {
	desc := `
	`
	sol := Solution{
		Title:  "最大正方形",
		Desc:   desc,
		Method: reflect.ValueOf(maximalSquare),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}}
	a.Output = []interface{}{4}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0221"] = sol
}
