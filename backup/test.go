package main

import (
	"fmt"
)

func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func IntMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	alen := len(nums1)
	blen := len(nums2)

	m := IntMin(alen, blen)
	n := IntMax(alen, blen)

	var i, j int = 0, 0
	var imin, imax int = 0, m - 1

	for true {
		// binary search
		i := (imin + imax) / 2
		j := (m+n+1)/2 - i

		// 满足条件退出
		//	1. j = (m+n+1)/2 - i
		//	2. max(L) <= min(R) L: left set R: right set
	}

	return 0
}
*/

func lengthOfLongestSubstringv2(s string) int {
	s_len := len(s)
	bmap := make([]int, 128)

	count := 0
	for i, j := 0, 0; j < s_len; j++ {
		idx := int(s[j])
		find := bmap[idx]
		if find > i {
			i = find
		}

		ans := j - i + 1
		if ans > count {
			count = ans
		}
		bmap[idx] = j + 1
	}

	return count
}

func lengthOfLongestSubstring(s string) int {
	bit_set := func(target []int, idx uint, set bool) {
		offset := idx / 32
		mask := 0x01 << (idx % 32)
		if !set {
			mask = ^mask
		}

		if set {
			target[offset] = target[offset] | mask
		} else {
			target[offset] = target[offset] & mask
		}
	}
	bit_get := func(target []int, idx uint) int {
		offset := idx / 32
		mask := 0x01 << (idx % 32)

		result := mask & target[offset]
		if result == mask {
			return 1
		}

		return 0
	}

	bit_map := make([]int, 4)
	reset_map := func() {
		bit_map[0] = 0
		bit_map[1] = 0
		bit_map[2] = 0
		bit_map[3] = 0
	}
	var bit_count int
	var max_bit_count int

	for idx := 0; idx < len(s); idx++ {
		char_int_val := uint(s[idx])
		has_char := bit_get(bit_map, char_int_val)
		if has_char == 0 {
			bit_set(bit_map, char_int_val, true)
			bit_count += 1
			if bit_count > max_bit_count {
				max_bit_count = bit_count
			}
		} else {
			reset_map()
			fmt.Println("---> ", bit_map)
			bit_set(bit_map, char_int_val, true)
			bit_count = 1
			if bit_count > max_bit_count {
				max_bit_count = bit_count
			}
		}

	}

	return max_bit_count
}

func main() {
	/*
		bit_set := func(target int, idx uint, set bool) int {
			mask := 0x01 << idx
			if !set {
				mask = ^mask
			}

			ret := target
			if set {
				ret = target | mask
			} else {
				ret = target & mask
			}

			return ret
		}

		test_int := 2
		ret_int := bit_set(test_int, 1, false)
		fmt.Println("位操作: ", test_int, " : ", ret_int)
	*/

	/*
		test_str := "dvdf"
		fmt.Println("结果: ", lengthOfLongestSubstringv2(test_str))
	*/

	/*
		nums1 := []int{1, 2, 6, 9, 12}
		nums2 := []int{5, 7, 8, 10, 11}
		ret := findMedianSortedArrays(nums1, nums2)
		fmt.Println("中位数是: ", ret)
	*/

	test_str := "abcdefghijklmn"
	fmt.Println("字符串长度: ", len(test_str), " # ", test_str[0:14])
	span := len(test_str)
	for i := 0; i < span-1; i++ {
		dd := span - i
		start := i
		fmt.Println(start, " -> ", dd)
		tmp_str := test_str[start : start+dd]
		fmt.Println(test_str)
		fmt.Println(tmp_str)
	}
}
