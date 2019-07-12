package solutions

import (
	"reflect"
)

func canJump(nums []int) bool {
	return false
}

func init() {
	desc := `
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
	`
	sol := Solution{
		Title:  "Jump Game",
		Desc:   desc,
		Method: reflect.ValueOf(canJump),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{[]int{2, 3, 1, 1, 4}}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{3, 2, 1, 0, 4}}
	a.Output = []interface{}{false}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0055"] = sol
}
