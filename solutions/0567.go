package solutions

import (
	"reflect"
)

func getAlphaValue(b byte) int {
	return int(b - 'a')
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func checkInclusion(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len > s2Len {
		return false
	}

	mask := make([]int, 26)
	slide := make([]int, 26)
	for i := 0; i < s1Len; i++ {
		mask[getAlphaValue(s1[i])]++
		slide[getAlphaValue(s2[i])]++
	}

	for i := 0; i < (s2Len - s1Len); i++ {
		if isEqual(mask, slide) {
			return true
		}
		e := i + s1Len
		slide[getAlphaValue(s2[i])]--
		slide[getAlphaValue(s2[e])]++
	}
	if isEqual(mask, slide) {
		return true
	}

	return false
}

func init() {
	desc := `
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.



Example 1:

Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input:s1= "ab" s2 = "eidboaoo"
Output: False


Note:

The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].

	`
	sol := Solution{
		Title:  "字符串的排列",
		Desc:   desc,
		Method: reflect.ValueOf(checkInclusion),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"", ""}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0567"] = sol
}
