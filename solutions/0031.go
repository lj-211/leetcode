package solutions

import (
	"reflect"
)

func nextPermutation(nums []int) {
	return
}

func init() {
	desc := `
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
	`
	sol := Solution{
		Title:  "Next Permutation",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	//a := TestCase{}
	//sol.Tests = append(sol.Tests, a)

	SolutionMap["0031"] = sol
}
