package solutions

import (
	"fmt"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func trap(height []int) int {
	size := len(height)
	if size <= 1 {
		return 0
	}

	dpLeft := make([]int, size)
	dpRight := make([]int, size)

	for i := 1; i < size; i++ {
		dpLeft[i] = ds.MaxInt(dpLeft[i-1], height[i-1])
	}
	for i := size - 2; i >= 0; i-- {
		dpRight[i] = ds.MaxInt(dpRight[i+1], height[i+1])
	}

	fmt.Println(dpLeft)
	fmt.Println(dpRight)

	cnt := 0
	for i := 1; i < size-1; i++ {
		minHeight := ds.MinInt(dpLeft[i], dpRight[i])
		if minHeight > height[i] {
			cnt += (minHeight - height[i])
		}
	}

	return cnt
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
		Title:  "Trapping Rain Water",
		Desc:   desc,
		Method: reflect.ValueOf(trap),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0042"] = sol
}
