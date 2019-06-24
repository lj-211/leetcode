package solutions

import (
	"reflect"
)

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	/*
		dp := make([]int, n)
		dp[0] = 1
		dp[1] = 2
		for i := 2; i < n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
	*/
	one := 2
	two := 1
	cur := 0
	for i := 2; i < n; i++ {
		cur = one + two
		two = one
		one = cur
	}
	return cur
}

func init() {
	desc := `
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
	`
	sol := Solution{
		Title:  "Climbing Stairs",
		Desc:   desc,
		Method: reflect.ValueOf(climbStairs),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{2}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0070"] = sol
}
