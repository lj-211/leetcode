package solutions

import (
	"fmt"
	"reflect"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	smap := make(map[string][]string)

	for _, v := range strs {
		intSlice := make([]int, 0)
		for _, s := range v {
			intSlice = append(intSlice, int(s))
		}
		sort.Ints(intSlice)
		isStr := fmt.Sprintf("%v", intSlice)
		smap[isStr] = append(smap[isStr], v)
	}
	ret := make([][]string, 0)
	for _, v := range smap {
		ret = append(ret, v)
	}
	return ret
}

func init() {
	desc := `
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
	`
	sol := Solution{
		Title:  "Group Anagrams",
		Desc:   desc,
		Method: reflect.ValueOf(groupAnagrams),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}
	a.Output = []interface{}{[]string{"ate", "eat", "tea"}, []string{"nat", "tan"}, []string{"bat"}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0049"] = sol
}
