package solutions

import (
	"reflect"
)

func jump(nums []int) int {
	return -1
}

func init() {
	desc := `
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:

Input: [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
    Jump 1 step from index 0 to 1, then 3 steps to the last index.
	`
	sol := Solution{
		Title:  "Jump Game II",
		Desc:   desc,
		Method: reflect.ValueOf(jump),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{[]int{2, 3, 1, 1, 4}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0045"] = sol
}
