package solutions

import (
	"reflect"
)

func isPalindrome(pstr string) bool {
	size := len(pstr)
	if size == 0 {
		return true
	}

	i := 0
	j := size - 1
	for i <= j {
		if pstr[i] != pstr[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func backtracking(s string, output *[][]string, cur []string) {
	size := len(s)
	if size == 0 {
		*output = append(*output, cur)
		return
	}
	newCur := make([]string, 0)
	newCur = append(newCur, cur...)
	newCur = append(newCur, s[0:1])
	backtracking(s[1:], output, newCur)

	for i := 2; i <= size; i++ {
		curStr := s[0:i]
		if isPalindrome(curStr) {
			newCur := make([]string, 0)
			newCur = append(newCur, cur...)
			newCur = append(newCur, curStr)

			backtracking(s[i:], output, newCur)
		}
	}
}

func partition(s string) [][]string {
	output := make([][]string, 0)
	backtracking(s, &output, []string{})
	return output
}

func init() {
	desc := `
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
	`
	sol := Solution{
		Title:  "Palindrome Partitioning",
		Desc:   desc,
		Method: reflect.ValueOf(partition),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"aab"}
	a.Output = []interface{}{[][]string{
		[]string{"aa", "b"},
		[]string{"a", "a", "b"},
	}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0131"] = sol
}
