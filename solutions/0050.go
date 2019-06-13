package solutions

import (
	"reflect"
)

func myPow(x float64, n int) float64 {
	return 0
}

func init() {
	desc := `
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
	`
	sol := Solution{
		Title:  "Pow(x, n)",
		Desc:   desc,
		Method: reflect.ValueOf(myPow),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{2, 10}
	a.Output = []interface{}{1024}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0050"] = sol
}
