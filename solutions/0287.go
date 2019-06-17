package solutions

import (
	"reflect"
)

func findDuplicate(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	p1 := nums[0]
	p2 := tortoise

	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}

	return p1
}

func init() {
	desc := `
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2
Example 2:

Input: [3,1,3,4,2]
Output: 3
Note:

You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
	`
	sol := Solution{
		Title:  "Find the Duplicate Number",
		Desc:   desc,
		Method: reflect.ValueOf(findDuplicate),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 3, 4, 2, 2}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0287"] = sol
}
