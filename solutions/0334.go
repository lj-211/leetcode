package solutions

import (
	"reflect"
)

func increasingTriplet(nums []int) bool {
	size := len(nums)
	if size <= 2 {
		return false
	}

	dpMin := make([]int, size)
	dpMin[0] = 0
	for i := 1; i < size; i++ {
		ival := nums[i]
		dpVal := nums[dpMin[i-1]]
		if ival > dpVal {
			dpMin[i] = dpMin[i-1]
		} else if ival <= dpVal {
			dpMin[i] = i
		}
	}
	dpMax := make([]int, size)
	dpMax[size-1] = size - 1
	for i := size - 2; i >= 0; i-- {
		ival := nums[i]
		dpVal := nums[dpMax[i+1]]
		if ival < dpVal {
			dpMax[i] = dpMax[i+1]
		} else if ival >= dpVal {
			dpMax[i] = i
		}
	}

	exist := false
	for i := 1; i < size-1; i++ {
		l := dpMin[i]
		r := dpMax[i]

		if i > l && i < r && nums[i] > nums[l] && nums[i] < nums[r] {
			exist = true
			break
		}
	}

	return exist
}

func init() {
	desc := `
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

数学表达式如下:

如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。

示例 1:

输入: [1,2,3,4,5]
输出: true
示例 2:

输入: [5,4,3,2,1]
输出: false
	`
	sol := Solution{
		Title:  "递增的三元子序列",
		Desc:   desc,
		Method: reflect.ValueOf(increasingTriplet),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 4, 5}}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0334"] = sol
}
