package solutions

import (
	"math"
	"sort"
)

/*
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	size := len(nums)
	if size < 3 {
		return -1
	}

	absFunc := func(a int) int {
		if a < 0 {
			return -1 * a
		}
		return a
	}

	mostCloseSum := math.MaxInt32
	closestSum := -1

	for i := 0; i < size; i++ {
		low := 0
		high := size - 1
		for low < high {
			if low == i {
				low++
				continue
			}
			if high == i {
				high--
				continue
			}

			sum := nums[i] + nums[low] + nums[high]
			delta := sum - target
			abs := absFunc(delta)
			if abs < mostCloseSum {
				mostCloseSum = abs
				closestSum = sum
			}

			if delta < 0 {
				low++
			} else if delta > 0 {
				high--
			} else {
				low++
				high--
			}
		}
	}

	return closestSum
}
