package solutions

import (
	"reflect"
)

func countSmaller(nums []int) []int {
	return []int{}
}

func init() {
	desc := `
You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

Example:

Input: [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
	`
	sol := Solution{
		Title:  "Count of Smaller Numbers After Self",
		Desc:   desc,
		Method: reflect.ValueOf(countSmaller),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{5, 2, 6, 1}}
	a.Output = []interface{}{[]int{2, 1, 1, 0}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0001"] = sol
}
