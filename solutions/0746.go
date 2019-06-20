package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	if size == 0 {
		return 0
	}
	f2 := cost[0]
	f1 := cost[1]
	f := 0
	for i := 2; i < len(cost); i++ {
		f = cost[i] + ds.MinInt(f1, f2)
		f2 = f1
		f1 = f
	}

	return ds.MinInt(f1, f2)
}

func init() {
	desc := `
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:

Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
Example 2:

Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
Note:

cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
	`
	sol := Solution{
		Title:  "Min Cost Climbing Stairs",
		Desc:   desc,
		Method: reflect.ValueOf(minCostClimbingStairs),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{10, 15, 20}}
	a.Output = []interface{}{15}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0746"] = sol
}
