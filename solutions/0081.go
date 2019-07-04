package solutions

import (
	"fmt"
	"reflect"
)

func searchII(nums []int, target int) bool {
	fmt.Println("searchII")
	size := len(nums)
	if size == 0 {
		return false
	}
	if size == 1 {
		if target == nums[0] {
			return true
		} else {
			return false
		}
	}

	l := 0
	r := size - 1
	for l <= r {
		lval := nums[l]
		rval := nums[r]
		m := (l + r) >> 1
		mval := nums[m]
		if mval == target {
			return true
		}
		if lval == mval && rval == mval {
			l++
			r--
			continue
		}

		if lval <= mval { // 左边是顺序递增
			if target >= lval && target < mval {
				r = m - 1
			} else {
				l = m + 1
			}
		} else { // 右边是顺序递增
			if target > mval && target <= rval {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

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

	/*
			a.Input = []interface{}{[]int{2, 5, 6, 0, 0, 1, 2}, 0}
			a.Output = []interface{}{true}
			sol.Tests = append(sol.Tests, a)

			a.Input = []interface{}{[]int{2, 5, 6, 0, 0, 1, 2}, 3}
			a.Output = []interface{}{false}
			sol.Tests = append(sol.Tests, a)

			a.Input = []interface{}{[]int{1, 1}, 1}
			a.Output = []interface{}{true}
			sol.Tests = append(sol.Tests, a)

		a.Input = []interface{}{[]int{3, 1}, 1}
		a.Output = []interface{}{true}
		sol.Tests = append(sol.Tests, a)
	*/

	a.Input = []interface{}{[]int{3, 1, 1}, 3}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 3, 1, 1, 1, 1}, 3}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{3, 1}, 1}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0081"] = sol
}
