package solutions

import (
	"fmt"
	"reflect"
)

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	alphas := make([]int, 0)
	amap := make(map[int]int)

	for _, v := range s {
		vint := int(v) - int('a')
		if c, ok := amap[vint]; ok {
			amap[vint] = c + 1
		} else {
			alphas = append(alphas, vint)
			amap[vint] = 1
		}
	}

	idx := -1
	for k, v := range alphas {
		c, _ := amap[v]
		if c == 1 {
			idx = k
			break
		}
	}

	return idx
}

func init() {
	desc := `
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:
s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
	`
	sol := Solution{
		Title:  "First Unique Character in a String",
		Desc:   desc,
		Method: reflect.ValueOf(firstUniqChar),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"leetcode"}
	a.Output = []interface{}{0}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0387"] = sol
}
