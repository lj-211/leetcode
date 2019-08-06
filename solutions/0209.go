package solutions

import (
	"math"
	"reflect"
)

func minSubArrayLen(s int, nums []int) int {
	notExist := 0
	size := len(nums)
	if size == 0 {
		return notExist
	}

	b := 0
	e := 0
	sum := 0
	minLen := math.MaxInt32

	expand := true

	for b < size && e < size {
		if e < b {
			e = b
			sum = nums[b]
		} else {
			if expand {
				sum += nums[e]
			} else {
				sum -= nums[b-1]
			}
		}
		if sum >= s {
			delta := e - b + 1
			if delta < minLen {
				minLen = delta
			}

			b++
			expand = false
		} else {
			e++
			expand = true
		}
	}
	if minLen == math.MaxInt32 {
		minLen = notExist
	}

	return minLen
}

func init() {
	desc := `
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例: 

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
	`
	sol := Solution{
		Title:  "长度最小子数组",
		Desc:   desc,
		Method: reflect.ValueOf(minSubArrayLen),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{7, []int{2, 3, 1, 2, 4, 3}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0209"] = sol
}
