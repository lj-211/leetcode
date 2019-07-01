package solutions

import (
	"reflect"
)

func removeElement(nums []int, val int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	s := 0
	for i := 0; i < size; i++ {
		if nums[i] != val {
			nums[s] = nums[i]
			s++
		}
	}

	return s
}

func init() {
	desc := `
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
	`
	sol := Solution{
		Title:  "Remove Element",
		Desc:   desc,
		Method: reflect.ValueOf(removeElement),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}
	a.Output = []interface{}{5}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0027"] = sol
}
