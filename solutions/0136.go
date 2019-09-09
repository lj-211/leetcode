package solutions

import (
	"reflect"
)

func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}

	return ret
}

func init() {
	desc := `
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
	`
	sol := Solution{
		Title:  "只出现一次的数字",
		Desc:   desc,
		Method: reflect.ValueOf(singleNumber),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 2, 1}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0136"] = sol
}
