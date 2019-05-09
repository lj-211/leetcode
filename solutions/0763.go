package solutions

import (
	"reflect"
)

func partitionLabels(S string) []int {
	last_idx := make(map[byte]int)
	slen := len(S)
	for i := 0; i < slen; i++ {
		last_idx[S[i]] = i
	}

	ret := make([]int, 0)

	var start int = 0
	var last int = 0
	for i := 0; i < slen; i++ {
		clast, _ := last_idx[S[i]]
		if clast > last {
			last = clast
		}
		if last == (slen - 1) {
			ret = append(ret, slen-start)
			break
		}

		if i == last {
			ret = append(ret, last-start+1)
			start = i + 1
			continue
		}
	}

	return ret
}

func init() {
	desc := `
A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

Example 1:
	Input: S = "ababcbacadefegdehijhklij"
	Output: [9,7,8]
	Explanation:
	The partition is "ababcbaca", "defegde", "hijhklij".
	This is a partition so that each letter appears in at most one part.
	A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
Note:
	S will have length in range [1, 500].
	S will consist of lowercase letters ('a' to 'z') only.
	`
	sol := Solution{
		Title:  "Partition Labels",
		Desc:   desc,
		Method: reflect.ValueOf(partitionLabels),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"ababcbacadefegdehijhklij"}
	a.Output = []interface{}{[]int{9, 7, 8}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0763"] = sol
}
