package solutions

import (
	"fmt"
	"reflect"
	"sort"
)

const (
	Start = iota
	End
)

type InfoTmp struct {
	Val  int
	Flag int
}

type Infos []InfoTmp

func (s Infos) Len() int      { return len(s) }
func (s Infos) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Infos) Less(i, j int) bool {
	if s[i].Val == s[j].Val {
		return s[i].Flag < s[j].Flag
	}
	return s[i].Val < s[j].Val
}

func merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size == 0 {
		return [][]int{}
	}
	if size == 1 {
		return intervals
	}
	var allInts Infos
	for i := 0; i < len(intervals); i++ {
		allInts = append(allInts, InfoTmp{
			Val:  intervals[i][0],
			Flag: Start,
		})
		allInts = append(allInts, InfoTmp{
			Val:  intervals[i][1],
			Flag: End,
		})
	}

	sort.Sort(allInts)
	fmt.Println(allInts)

	ret := make([][]int, 0)

	stack := 1
	s := 0
	for i := 1; i < len(allInts); i++ {
		v := allInts[i]
		if v.Flag == Start {
			stack++
		} else {
			stack--
			if stack == 0 {
				one := make([]int, 2)
				one[0] = allInts[s].Val
				one[1] = allInts[i].Val
				ret = append(ret, one)
				s = i + 1
			}
		}
	}

	return ret
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
	a.Input = []interface{}{[][]int{[]int{1, 4}, []int{4, 5}}}
	a.Output = []interface{}{[]int{1, 5}}
	sol.Tests = append(sol.Tests, a)
	/*
		a.Input = []interface{}{[][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}}
		a.Output = []interface{}{[][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}}}
		sol.Tests = append(sol.Tests, a)
	*/

	SolutionMap["0056"] = sol
}
