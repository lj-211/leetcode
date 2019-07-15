package solutions

import (
	"reflect"
)

func letterCombinations(digits string) []string {
	return []string{}
}

func init() {
	desc := `
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
	`
	sol := Solution{
		Title:  "Letter Combinations of a Phone Number",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"23"}
	a.Output = []interface{}{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0017"] = sol
}
