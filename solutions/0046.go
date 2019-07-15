package solutions

import (
	"reflect"
)

func permute(nums []int) [][]int {
	return [][]int{[]int{}}
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
	//a := TestCase{}
	//sol.Tests = append(sol.Tests, a)

	SolutionMap["0046"] = sol
}
