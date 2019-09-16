package solutions

import (
	"reflect"
)

func MaxInt32(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if size < k || size == 0 || k == 0 {
		return []int{}
	}

	left := make([]int, size)
	for i := 0; i < size; i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = MaxInt32(nums[i], left[i-1])
		}
	}

	right := make([]int, size)
	right[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		if (i % k) == (k - 1) {
			right[i] = nums[i]
		} else {
			right[i] = MaxInt32(nums[i], right[i+1])
		}
	}

	retSize := size - k + 1
	ret := make([]int, retSize)
	for j := k - 1; j < size; j++ {
		i := j - k + 1
		ret[i] = MaxInt32(left[j], right[i])
	}

	return ret
}

func init() {
	desc := `
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

 

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 

提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
	`
	sol := Solution{
		Title:  "滑动窗口最大值",
		Desc:   desc,
		Method: reflect.ValueOf(maxSlidingWindow),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 3, -1, -3, 5, 3, 6, 7}}
	a.Output = []interface{}{[]int{3, 3, 5, 5, 6, 7}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0239"] = sol
}
