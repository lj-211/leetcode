package solutions

import (
	"reflect"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	size := len(nums)

	cnt := 0
	/*
		for i := 0; i < size; i++ {
			cnt++
			muls := nums[i]
			for j := i + 1; j < size; j++ {
				muls *= nums[j]
				if muls >= k {
					break
				} else {
					cnt++
					fmt.Println("i-j: ", i, j)
				}
			}
		}
	*/

	l := 0
	muls := 1
	for i := 0; i < size; i++ {
		muls *= nums[i]
		for muls >= k && l <= i {
			muls = muls / nums[l]
			l++
		}
		if l <= i {
			cnt += (i - l + 1)
		}
	}

	return cnt
}

func init() {
	desc := `
给定一个正整数数组 nums。

找出该数组内乘积小于 k 的连续的子数组的个数。

示例 1:

输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8个乘积小于100的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
说明:

0 < nums.length <= 50000
0 < nums[i] < 1000
0 <= k < 10^6
	`
	sol := Solution{
		Title:  "乘积小于K的子数组",
		Desc:   desc,
		Method: reflect.ValueOf(numSubarrayProductLessThanK),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{10, 5, 2, 6}, 100}
	a.Output = []interface{}{100}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0713"] = sol
}
