package solutions

import (
	"reflect"
)

func productExceptSelf(nums []int) []int {
	size := len(nums)
	if size == 0 || size == 1 {
		return []int{}
	}
	dpLeft := make([]int, size)
	dpLeft[0] = 1
	for i := 1; i < size; i++ {
		dpLeft[i] = dpLeft[i-1] * nums[i-1]
	}

	dpRight := make([]int, size)
	dpRight[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		dpRight[i] = dpRight[i+1] * nums[i+1]
	}

	ret := make([]int, size)
	for i := 0; i < size; i++ {
		ret[i] = dpLeft[i] * dpRight[i]
	}

	return ret
}

func init() {
	desc := `
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
	`
	sol := Solution{
		Title:  "除自身以外数组的乘积",
		Desc:   desc,
		Method: reflect.ValueOf(productExceptSelf),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 4}}
	a.Output = []interface{}{[]int{24, 12, 8, 6}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0001"] = sol
}
