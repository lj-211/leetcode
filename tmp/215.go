/*
 * 215. Kth Largest Element in an Array
 * Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
 *
 * Example 1:
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 *
 * Example 2:
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 */
package solutions

import (
	"fmt"
)

func partition(nums []int, s int, e int) int {
	pivot := nums[e]

	j := s
	i := s - 1
	for j <= e-1 {
		if nums[j] <= pivot {
			i++
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
		j++
	}
	i++
	tmp := nums[i]
	nums[i] = nums[e]
	nums[e] = tmp

	fmt.Println("交换: ", i, " ", e)
	fmt.Println("调整后的数组: ", nums)

	return i
}

func quick(nums []int, s int, e int, dst_idx int) int {
	if s < e {
		idx := partition(nums, s, e)
		if idx == dst_idx {
			return nums[dst_idx]
		}
		if dst_idx < idx {
			fmt.Println("左边: ", s, " - ", idx-1)
			return quick(nums, s, idx-1, dst_idx)
		} else {
			fmt.Println("右边: ", idx+1, " - ", e)
			return quick(nums, idx+1, e, dst_idx)
		}
	}

	return 0
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	dst_idx := len(nums) - k
	quick(nums, 0, len(nums)-1, dst_idx)

	return nums[dst_idx]
}

func main() {
	//input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	input := []int{3, 2, 1, 5, 6, 4}
	//input := []int{1}
	//input := []int{3, 1, 2, 4}
	fmt.Println("Kth Largest Element in an Array")
	fmt.Println("output: ", findKthLargest(input, 2))
}
