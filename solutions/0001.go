package solutions

import (
	"reflect"
)

func twoSum(nums []int, target int) []int {
	int_map := make(map[int]int)
	for idx, v := range nums {
		int_map[v] = idx
	}

	ret := make([]int, 0)
	for idx, v := range nums {
		if tidx, ok := int_map[target-v]; ok {
			if tidx != idx {
				ret = append(ret, idx)
				ret = append(ret, tidx)
				break
			}
		}
	}

	return ret
}

func init() {
	sol := Solution{
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0001"] = sol
}
