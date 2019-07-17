package solutions

import (
	"reflect"
)

func sortColors(nums []int) []int {
	size := len(nums)
	start := 0
	end := size - 1
	fmt.Println(nums)
	for i := 0; i <= end; {
		if nums[i] == 0 {
			nums[start], nums[i] = nums[i], nums[start]
			start++
			i++
		} else if nums[i] == 2 {
			nums[end], nums[i] = nums[i], nums[end]
			end--
		} else {
			i++
		}
	}

	return nums
}

func init() {
	desc := `
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
	`
	sol := Solution{
		Title:  "Sort Colors",
		Desc:   desc,
		Method: reflect.ValueOf(sortColors),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	/*
		a.Input = []interface{}{[]int{2, 0, 2, 1, 1, 0}}
		a.Output = []interface{}{[]int{0, 0, 1, 1, 2, 2}}
		sol.Tests = append(sol.Tests, a)
	*/
	a.Input = []interface{}{[]int{2, 0, 1}}
	a.Output = []interface{}{[]int{0, 1, 2}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0075"] = sol
}
