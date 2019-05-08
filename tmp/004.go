/*
 * Median of Two Sorted Arrays
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 * You may assume nums1 and nums2 cannot be both empty.
 * Example 1:
 * 	nums1 = [1, 3]
 * 	nums2 = [2]
 *
 * 	The median is 2.0
 * Example 2:
 * 	nums1 = [1, 2]
 * 	nums2 = [3, 4]
 *
 * 	The median is (2 + 3)/2 = 2.5
 */
package solutions

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	IntMin := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	IntMax := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	alen := len(nums1)
	blen := len(nums2)

	var m, n int = 0, 0
	var marr []int
	var narr []int

	if alen > blen {
		m = blen
		marr = nums2
		n = alen
		narr = nums1
	} else {
		m = alen
		marr = nums1
		n = blen
		narr = nums2
	}

	var i, j int = 0, 0
	var imin, imax int = 0, m
	var ret float64

	for imin <= imax {
		// binary search
		i = (imin + imax) / 2
		j = (m+n+1)/2 - i

		// 满足条件退出
		//	1. j = (m+n+1)/2 - i
		//	2. max(L) <= min(R) L: left set R: right set
		if i < imax && narr[j-1] > marr[i] {
			imin = i + 1
		} else if i > imin && marr[i-1] > narr[j] {
			imax = i - 1
		} else {
			var maxleft int = 0
			var maxright int = 0

			for loop := true; loop; loop = false {
				if i == 0 {
					maxleft = narr[j-1]
				} else if j == 0 {
					maxleft = marr[i-1]
				} else {
					maxleft = IntMax(marr[i-1], narr[j-1])
				}
				if (alen+blen)%2 == 1 {
					ret = float64(maxleft)
					fmt.Println("maxleft: ", maxleft)
					break
				}

				if i == m {
					maxright = narr[j]
				} else if j == n {
					maxright = marr[i]
				} else {
					maxright = IntMin(marr[i], narr[j])
				}

				ret = float64(maxleft+maxright) / float64(2)
			}
			break
		}
	}

	return ret
}

func main() {
	// 1 2 5 6 7 8 9 10 11 12
	//nums1 := []int{1, 2, 6, 9}
	//nums2 := []int{5, 7, 8, 10}
	nums1 := []int{-2, -1}
	nums2 := []int{3}
	ret := findMedianSortedArrays(nums1, nums2)
	fmt.Println("中位数是: ", ret)
}
