package solutions

import (
	"reflect"
)

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = true
	}

	maxLen := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		clen := 1
		_, ok := hashMap[v-1]
		if !ok {
			step := v + 1
			for {
				_, ok := hashMap[step]
				if ok {
					step++
					clen++
				} else {
					break
				}
			}
		}

		if clen > maxLen {
			maxLen = clen
		}
	}

	return maxLen
}

func init() {
	desc := `
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
	`
	sol := Solution{
		Title:  "最长连续序列",
		Desc:   desc,
		Method: reflect.ValueOf(longestConsecutive),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0128"] = sol
}
