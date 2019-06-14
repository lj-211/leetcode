package solutions

import (
	//"fmt"
	"reflect"
)

func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i-1][j-1] * 2 + 2
	// dp[0][0] = 0 dp[1][1] = 1

	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// dp[0][1] = 1 dp[1][0] = 1
	/*
		dp := make([][]int, m)
		for i := 0; i < m; i++ {
			dp[i] = make([]int, n)
		}
		for i := 1; i < m; i++ {
			dp[i][0] = 1
		}
		for j := 1; j < n; j++ {
			dp[0][j] = 1
		}

		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	*/
	curj := make([]int, 0)
	for i := 0; i < n; i++ {
		curj = append(curj, 1)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			curj[j] += curj[j-1]
		}
	}

	return curj[n-1]
}

func init() {
	desc := `
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right

Example 2:
Input: m = 7, n = 3
Output: 28
	`
	sol := Solution{
		Title:  "Unique Paths",
		Desc:   desc,
		Method: reflect.ValueOf(uniquePaths),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{3, 2}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0062"] = sol
}
