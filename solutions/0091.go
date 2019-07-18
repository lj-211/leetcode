package solutions

import (
	"fmt"
	"reflect"
	//"strconv"
)

const MinAlpha int = 1
const MaxAlpha int = 26

// 本题的解法是backtracking
// note: dp的解法效率上更高一些

func _convertSpecialStrToInt(s string) int {
	h := int(s[0]) - int('1') + 1
	l := int(s[1]) - int('1') + 1
	return h*10 + l
}

func decodeWay(s string, cnt *int) {
	size := len(s)
	if size == 0 {
		(*cnt)++
		return
	}

	if s[0] == '0' {
		return
	}

	decodeWay(s[1:], cnt)

	if size >= 2 {
		two := s[0:2]

		//tval, _ := strconv.Atoi(two)
		tval := _convertSpecialStrToInt(two)
		fmt.Println("-> ", tval)
		if tval <= MaxAlpha {
			decodeWay(s[2:], cnt)
		}
	}
}

func numDecodings(s string) int {
	cnt := 0

	decodeWay(s, &cnt)

	return cnt
}

func init() {
	desc := `
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
	`
	sol := Solution{
		Title:  "Decode Ways",
		Desc:   desc,
		Method: reflect.ValueOf(numDecodings),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	/*
		a.Input = []interface{}{"226"}
		a.Output = []interface{}{3}
		sol.Tests = append(sol.Tests, a)
	*/
	a.Input = []interface{}{"12"}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	/*
		a.Input = []interface{}{"30"}
		a.Output = []interface{}{0}
		sol.Tests = append(sol.Tests, a)
	*/

	SolutionMap["0091"] = sol
}
