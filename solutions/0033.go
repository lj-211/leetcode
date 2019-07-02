package solutions

import (
	"fmt"
	"reflect"
)

func search(nums []int, target int) int {
	size := len(nums)
	// 这里默认size >= 2
	if size == 0 {
		return -1
	}
	if size == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	pivot := nums[size-1]
	l := 0
	r := size - 2
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= pivot {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	dst := l
	if target > pivot {
		l = 0
		r = dst - 1
	} else {
		l = dst
		r = size - 1
	}

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			fmt.Println("返回: ", mid)
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
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
		Method: reflect.ValueOf(search),
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
