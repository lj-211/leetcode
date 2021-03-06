package solutions

import (
	"reflect"
)

func maxSubArray(nums []int) int {
	accu := nums[0]
	max_sub := nums[0]
	for i := 1; i < len(nums); i++ {
		ival := nums[i]
		if accu < 0 {
			accu = ival
		} else {
			accu += ival
		}
		if max_sub < accu {
			max_sub = accu
		}
	}

	return max_sub
}

func init() {
	desc := `
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
	`
	sol := Solution{
		Title:  "Maximum Subarray",
		Desc:   desc,
		Method: reflect.ValueOf(maxSubArray),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{[]int{-1}}
	a.Output = []interface{}{-1}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{[]int{-2, 1}}
	a.Output = []interface{}{1}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0053"] = sol
}
