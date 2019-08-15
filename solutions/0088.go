package solutions

import (
	"math"
	"reflect"
)

func merge80(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for i := m - 1; i >= 0; i-- {
		nums1[i+n], nums1[i] = nums1[i], nums1[i+n]
	}

	mcursor := n
	ncursor := 0
	cursor := 0
	for cursor < (m + n) {
		mv := math.MaxInt32
		if mcursor < (m + n) {
			mv = nums1[mcursor]
		}
		nv := math.MaxInt32
		if ncursor < n {
			nv = nums2[ncursor]
		}

		if mv <= nv {
			nums1[cursor] = mv
			mcursor++
		} else {
			nums1[cursor] = nv
			ncursor++
		}

		cursor++
	}
}

func init() {
	desc := `
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
	`
	sol := Solution{
		Title:  "合并两个有序数组",
		Desc:   desc,
		Method: reflect.ValueOf(merge80),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0080"] = sol
}
