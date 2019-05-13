package solutions

import (
	"reflect"
	"strconv"
)

func addBinary(a string, b string) string {
	var i, j int = len(a) - 1, len(b) - 1
	var step int = 0
	ret := ""
	for true {
		aint := 0
		bint := 0
		if i >= 0 {
			aint = int(a[i]) - int('0')
		}
		if j >= 0 {
			bint = int(b[j]) - int('0')
		}
		sum := aint + bint + step
		if sum >= 2 {
			step = 1
		} else {
			step = 0
		}
		ret = strconv.FormatInt(int64(sum%2), 10) + ret

		i--
		j--
		if i < 0 && j < 0 {
			if step > 0 {
				ret = "1" + ret
			}
			break
		}
	}

	return ret
}

func init() {
	desc := `
Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"
	`
	sol := Solution{
		Title:  "Add Binary",
		Desc:   desc,
		Method: reflect.ValueOf(addBinary),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"1010", "1011"}
	a.Output = []interface{}{"10101"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0067"] = sol
}
