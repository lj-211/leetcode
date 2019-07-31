package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func maxProduct(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	min := nums[0]
	max := nums[0]
	maxP := max

	for i := 1; i < size; i++ {
		v := nums[i]

		if v > 0 {
			max = ds.MaxInt(max*v, v)
			min = ds.MinInt(min*v, v)
		} else {
			omax := max
			max = ds.MaxInt(min*v, v)
			min = ds.MinInt(omax*v, v)
		}
		if max > maxP {
			maxP = max
		}
	}

	return maxP
}

func init() {
	desc := `
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
	`
	sol := Solution{
		Title:  "Maximum Product Subarray",
		Desc:   desc,
		Method: reflect.ValueOf(maxProduct),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 3, -2, 4}}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0001"] = sol
}
