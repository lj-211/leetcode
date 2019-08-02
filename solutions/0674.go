package solutions

import (
	"reflect"
)

func findLengthOfLCIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	maxCnt := 1
	accut := 1
	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			accut += 1
			if accut > maxCnt {
				maxCnt = accut
			}
		} else {
			accut = 1
		}
	}

	return maxCnt
}

func init() {
	desc := `
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

示例 1:

输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
示例 2:

输入: [2,2,2,2,2]
输出: 1
解释: 最长连续递增序列是 [2], 长度为1。
注意：数组长度不会超过10000。
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	`
	sol := Solution{
		Title:  "最长连续递增序列",
		Desc:   desc,
		Method: reflect.ValueOf(findLengthOfLCIS),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 2, 2, 2}}
	a.Output = []interface{}{1}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0674"] = sol
}
