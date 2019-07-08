package solutions

import (
	"reflect"
)

func findPeakElement(nums []int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	l := 0
	r := size - 1

	for l <= r {
		m := (l + r) / 2
		if (m + 1) >= size {
			break
		}
		mval := nums[m]
		mpval := nums[m+1]
		if mval < mpval {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}

func init() {
	desc := `
A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5
Explanation: Your function can return either index number 1 where the peak element is 2,
             or index number 5 where the peak element is 6.
Note:

Your solution should be in logarithmic complexity.
	`
	sol := Solution{
		Title:  "Find Peak Element",
		Desc:   desc,
		Method: reflect.ValueOf(findPeakElement),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 1}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0162"] = sol
}
