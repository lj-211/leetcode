package solutions

import (
	"reflect"
)

func searchII(nums []int, target int) bool {
	return false
}

func init() {
	desc := `
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
Follow up:

This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?
	`
	sol := Solution{
		Title:  "Search in Rotated Sorted Array II",
		Desc:   desc,
		Method: reflect.ValueOf(searchII),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{[]int{2, 5, 6, 0, 0, 1, 2}, 0}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{2, 5, 6, 0, 0, 1, 2}, 3}
	a.Output = []interface{}{false}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0081"] = sol
}
