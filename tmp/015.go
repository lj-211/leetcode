/*
 * 15. 3Sum
 * Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
 *
 * Note:
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * A solution set is:
 * [
 *   [-1, 0, 1],
 *   [-1, -1, 2]
 * ]
 */
package solutions

import (
	"fmt"
	"sort"
)

func IntMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func IntMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func step(i int, j int, h int, nums []int) {
	min := IntMin(i, j)
	max := IntMax(i, j)

	if h <= max {
		return
	}

	sum := nums[i] + nums[j] + nums[h]
	if sum > 0 {
		step(min, max, h-1, nums)
	} else {
		if sum == 0 {
			fmt.Println(nums[i], " ", nums[j], " ", nums[h])
			//step(min, max+1, h-1, nums)
			step(max, max+1, h-1, nums)
		} else {
			//step(min, max+1, h, nums)
			step(max, max+1, h, nums)
		}
	}
}

func threeSum(nums []int) [][]int {
	num_cnt := len(nums)
	if num_cnt == 0 {
		return [][]int{}
	}
	sort.Ints(nums)

	fmt.Println(nums)
	var low_i, low_j = 0, 1
	var high = len(nums) - 1
	step(low_i, low_j, high, nums)

	return nil
}

func main() {
	fmt.Println("3Sum")
	input := []int{-1, 0, 1, 2, -1, -4}
	threeSum(input)
}
