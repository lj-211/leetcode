package solutions

import (
	"reflect"
)

func twoSum(nums []int, target int) []int {
	int_map := make(map[int]int)
	for idx, v := range nums {
		int_map[v] = idx
	}

	ret := make([]int, 0)
	for idx, v := range nums {
		if tidx, ok := int_map[target-v]; ok {
			if tidx != idx {
				ret = append(ret, idx)
				ret = append(ret, tidx)
				break
			}
		}
	}

	return ret
}

func init() {
	desc := `
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
	`
	sol := Solution{
		Title:  "Two Sum",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0001"] = sol
}
