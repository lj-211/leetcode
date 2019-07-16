package solutions

import (
	"reflect"
)

func merge(intervals [][]int) [][]int {
	return [][]int{}
}

func init() {
	desc := `
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
	`
	sol := Solution{
		Title:  "Merge Intervals",
		Desc:   desc,
		Method: reflect.ValueOf(merge),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 4}, []int{4, 5}}
	a.Output = []interface{}{[]int{1, 5}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0056"] = sol
}
