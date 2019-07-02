package solutions

import (
	"reflect"
)

func search(nums []int, target int) int {
	return 0
}

func init() {
	desc := `
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
	`
	sol := Solution{
		Title:  "Search in Rotated Sorted Array",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 0}
	a.Output = []interface{}{4}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 3}
	a.Output = []interface{}{-1}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0033"] = sol
}
