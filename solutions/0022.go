package solutions

import (
	"reflect"
)

func trace(now string, open int, left int, output *[]string) {
	if left == 0 && open == 0 {
		*output = append(*output, now)
	}
	if left > 0 {
		trace(now+"(", open+1, left-1, output)
	}
	if open > 0 {
		trace(now+")", open-1, left, output)
	}
}

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	trace("", 0, n, &ret)

	return ret
}

func init() {
	desc := `
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
	`
	sol := Solution{
		Title:  "Generate Parentheses",
		Desc:   desc,
		Method: reflect.ValueOf(generateParenthesis),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{3}
	a.Output = []interface{}{[]string{"((()))", "(()())", "(())()", "()(())", "()()()"}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0022"] = sol
}
