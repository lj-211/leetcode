package solutions

import (
	"reflect"
)

func trap(height []int) int {
	return 0
}

func init() {
	desc := `
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
	`
	sol := Solution{
		Title:  "Two Sum",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0042"] = sol
}
