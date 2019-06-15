package solutions

import (
	"reflect"
)

func searchRange(nums []int, target int) []int {
	size := len(nums)
	left := 0
	right := size - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	ret := make([]int, 0)
	if left < size && nums[left] == target {
		ret = append(ret, left)
	} else {
		ret = append(ret, -1)
	}

	left = 0
	right = size - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		ret = append(ret, right)
	} else {
		ret = append(ret, -1)
	}

	return ret
}

func init() {
	desc := `
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
	`
	sol := Solution{
		Title:  "Find First and Last Position of Element in Sorted Array",
		Desc:   desc,
		Method: reflect.ValueOf(searchRange),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{5, 7, 7, 8, 8, 10}, 8}
	a.Output = []interface{}{[]int{3, 4}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0034"] = sol
}
