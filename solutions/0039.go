package solutions

import (
	"sort"
)

/*
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 *   [7],
 *     [2,2,3]
 * 	]
 * 	示例 2:
 *
 * 	输入: candidates = [2,3,5], target = 8,
 * 	所求解集为:
 * 	[
 * 	  [2,2,2,2],
 * 	  [2,3,3],
 * 	  [3,5]
 * 	]
 */

// dp[7] = 7 + (6 + dp[1]) + (3 + dp[4]) + (2 + dp[5])
func backtracking0039(candidates []int, arr *[]int, sum int, target int, out *[][]int) {
	if sum == target {
		newArr := make([]int, len(*arr))
		copy(newArr, *arr)
		*out = append(*out, newArr)
		return
	}
	if sum > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		v := candidates[i]
		asize := len(*arr)
		if asize > 0 && (*arr)[asize-1] < v {
			continue
		}
		*arr = append(*arr, v)
		backtracking0039(candidates, arr, sum+v, target, out)
		*arr = (*arr)[0 : len(*arr)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	arrList := make([]int, 0)
	sum := 0
	out := make([][]int, 0)
	backtracking0039(candidates, &arrList, sum, target, &out)

	return out
}
