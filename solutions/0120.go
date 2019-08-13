package solutions

import (
	"reflect"
)

func MinInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func minimumTotal(triangle [][]int) int {
	size := len(triangle)
	if size == 0 {
		return 0
	}

	dp := make([]int, size)
	dp[0] = triangle[0][0]

	for i := 1; i < size; i++ {
		line := triangle[i]
		for j := len(line) - 1; j >= 0; j-- {
			min := 0
			if j == 0 {
				min = dp[j]
			} else if j == len(line)-1 {
				min = dp[j-1]
			} else {
				min = MinInt(dp[j-1], dp[j])
			}
			dp[j] = line[j] + min
		}
	}

	min := dp[0]
	for i := 1; i < size; i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}

	return min
}

func init() {
	desc := `
	`
	sol := Solution{
		Title:  "三角形最小路径和",
		Desc:   desc,
		Method: reflect.ValueOf(minimumTotal),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0120"] = sol
}
