package solutions

import (
	"reflect"
)

func moveZeroes(nums []int) {
	size := len(nums)
	if size == 0 {
		return
	}
	valid := -1

	for i := 0; i < size; i++ {
		v := nums[i]
		if v != 0 {
			valid++
			nums[valid], nums[i] = nums[i], nums[valid]
		}
	}
}

func init() {
	desc := `
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
	`
	sol := Solution{
		Title:  "移动零",
		Desc:   desc,
		Method: reflect.ValueOf(moveZeroes),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{0, 1, 0, 3, 12}}
	a.Output = []interface{}{[]int{1, 3, 12, 0, 0}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0283"] = sol
}
