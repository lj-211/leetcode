package solutions

import (
	"reflect"
)

func bt(nums []int, mask map[int]bool, one []int, result *[][]int) {
	if len(nums) == len(one) {
		dst := make([]int, len(one))
		copy(dst, one)
		*result = append(*result, dst)
		return
	}
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if _, ok := mask[v]; ok {
			continue
		}

		one = append(one, v)
		mask[v] = true
		bt(nums, mask, one, result)
		one = one[0 : len(one)-1]
		delete(mask, v)
	}
}

func permute(nums []int) [][]int {
	mask := make(map[int]bool)
	one := make([]int, 0)
	result := make([][]int, 0)

	bt(nums, mask, one, &result)

	return result
}

func init() {
	desc := `
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
	`
	sol := Solution{
		Title:  "Permutations",
		Desc:   desc,
		Method: reflect.ValueOf(permute),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{5, 4, 6, 2}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0046"] = sol
}
