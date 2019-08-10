package solutions

import (
	"fmt"
	"reflect"
)

func lengthOfLISDP(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	dp := make([]int, size)
	dp[0] = 1
	maxCnt := 0

	for i := 0; i < size; i++ {
		newDp := 1
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				newDp = dp[j] + 1
			} else {
				newDp = 1
			}
			if newDp > max {
				max = newDp
			}
		}
		if max > maxCnt {
			maxCnt = max
		}
		dp[i] = max
	}

	return maxCnt
}

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	dp := make([]int, size)
	dp[0] = 1
	minList := make([]int, size)
	minList[0] = nums[0]
	rc := 0

	for i := 0; i < size; i++ {
		l := 0
		r := rc
		v := nums[i]
		for l <= r {
			m := (l + r) / 2
			mval := minList[m]
			if i == 3 {
				fmt.Println("5: ", l, r, m, v, mval)
			}
			if v > mval {
				l = m + 1
			} else {
				r = m - 1
			}
			if i == 3 {
				fmt.Println("退出后: ", l)
			}
		}
		minList[l] = v
		if l > rc {
			rc = l
		}

		fmt.Println(minList, rc, l, r)
	}

	return rc + 1
}

func init() {
	desc := `
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
	`
	sol := Solution{
		Title:  "最长上升子序列",
		Desc:   desc,
		Method: reflect.ValueOf(lengthOfLIS),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	/*
		a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
		a.Output = []interface{}{[]int{0, 1}}
		sol.Tests = append(sol.Tests, a)
	*/

	a.Input = []interface{}{[]int{10, 9, 2, 5, 3, 7, 101, 18}}
	a.Output = []interface{}{4}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0300"] = sol
}
