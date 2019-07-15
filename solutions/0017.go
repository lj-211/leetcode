package solutions

import (
	"reflect"
)

var DigitMap map[byte][]string = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func recurComb(digits string) []string {
	size := len(digits)
	if size == 0 {
		return []string{}
	}

	d := digits[0]
	dslist := DigitMap[d]
	ret := make([]string, 0)
	comList := recurComb(digits[1:])
	for _, v := range dslist {
		if len(comList) > 0 {
			for _, v2 := range comList {
				ret = append(ret, v+v2)
			}
		} else {
			ret = append(ret, v)
		}
	}

	return ret
}

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
		Method: reflect.ValueOf(recurComb),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"23"}
	a.Output = []interface{}{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0017"] = sol
}
