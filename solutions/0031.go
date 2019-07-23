package solutions

import (
	"reflect"
)

func reverse(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	i, j := 0, size-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	for i := size - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			pivotIdx := i - 1

			j := size - 1
			for ; j >= i; j-- {
				if nums[j] > nums[pivotIdx] {
					break
				}
			}
			nums[pivotIdx], nums[j] = nums[j], nums[pivotIdx]

			reverse(nums[i:])

			return
		}
	}

	reverse(nums)

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
