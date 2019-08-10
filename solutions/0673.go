package solutions

import (
	"fmt"
	"reflect"
)

func findNumberOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	dp := make([]int, size)
	dp[0] = 1
	dpCnt := make([]int, size)
	dpCnt[0] = 1

	for i := 0; i < size; i++ {
		max := 1
		cnt := 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			newDp := dp[j] + 1
			if newDp > max {
				max = newDp
				cnt = dpCnt[j]
			} else if newDp == max {
				cnt += dpCnt[j]
			}
		}
		if false && max == 1 {
			cnt++
		}
		dp[i] = max
		dpCnt[i] = cnt
	}

	fmt.Println(dp)
	fmt.Println(dpCnt)

	max := 0
	maxCnt := 1
	for i := 0; i < size; i++ {
		v := dp[i]
		if v > max {
			max = v
			maxCnt = dpCnt[i]
			fmt.Println("设置: ", max, maxCnt, i, "-", dpCnt)
		} else if v == max {
			maxCnt += dpCnt[i]
		}
	}
	fmt.Println("max: ", max)
	fmt.Println("max: ", maxCnt)

	return maxCnt
}

func init() {
	desc := `
给定一个未排序的整数数组，找到最长递增子序列的个数。

示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:

输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。
	`
	sol := Solution{
		Title:  "最长递增子序列的个数",
		Desc:   desc,
		Method: reflect.ValueOf(findNumberOfLIS),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{[]int{1, 3, 5, 4, 7}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{2, 2, 2, 2, 2}}
	a.Output = []interface{}{5}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 2, 4, 3, 5, 4, 7, 2}}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0673"] = sol
}
