package solutions

import (
	"fmt"
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	size := len(nums)

	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		h := i + 1
		t := size - 1
		val := nums[i]
		for h < t {
			sum := val + nums[h] + nums[t]
			if sum > 0 {
				t--
			} else if sum < 0 {
				h++
			} else {
				arr := make([]int, 0)
				hval := nums[h]
				tval := nums[t]
				arr = append(arr, nums[i], nums[h], nums[t])
				ret = append(ret, arr)
				for h < t {
					h++
					if nums[h] != hval {
						break
					}
				}
				for h < t {
					t--
					if nums[t] != tval {
						break
					}
				}
			}
		}
	}

	fmt.Println(ret)
	return ret
}

func init() {
	desc := `
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
	`
	sol := Solution{
		Title:  "3Sum",
		Desc:   desc,
		Method: reflect.ValueOf(threeSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{-1, 0, 1, 2, -1, -4}}
	a.Output = []interface{}{}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0015"] = sol
}
