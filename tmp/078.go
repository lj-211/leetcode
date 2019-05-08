/*
 * 78. Subsets
 * Given a set of distinct integers, nums, return all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 *   [3],
 *   [1],
 *   [2],
 *   [1,2,3],
 *   [1,3],
 *   [2,3],
 *   [1,2],
 *   []
 * ]
 */
package solutions

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		single := nums[i]
		tmp := make([][]int, 0)
		for j := 0; j < len(ret); j++ {
			//new_arr := append(ret[j], single)
			new_arr := make([]int, 0)
			for t := 0; t < len(ret[j]); t++ {
				new_arr = append(new_arr, ret[j][t])
			}
			new_arr = append(new_arr, single)
			tmp = append(tmp, new_arr)
		}
		single_arr := make([]int, 0)
		single_arr = append(single_arr, single)
		ret = append(ret, single_arr)
		ret = append(ret, tmp...)
	}
	ret = append(ret, make([]int, 0))

	return ret
}

func main() {
	fmt.Println("全排列")
	//input := []int{1, 2, 3}
	input := []int{9, 0, 3, 5, 7}
	output := subsets(input)
	fmt.Println(output)
}
