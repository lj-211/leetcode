package solutions

import (
	"math"
	"reflect"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if dividend == math.MinInt32 || dividend == -1 {
		return math.MaxInt32
	}

	sign := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}

	// abs
	if dividend < 0 {
		dividend *= -1
	}
	if divisor < 0 {
		divisor *= -1
	}

	dvd := dividend
	ret := 0
	for dvd >= divisor {
		m := divisor
		var acc uint = 0
		for dvd >= (m << 1) {
			acc++
			m <<= 1
		}
		dvd = dvd - m
		ret += (1 << acc)
	}

	return ret * sign
}

func init() {
	desc := `
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
	`
	sol := Solution{
		Title:  "Divide Two Integers",
		Desc:   desc,
		Method: reflect.ValueOf(divide),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	/*
		a.Input = []interface{}{10, 3}
		a.Output = []interface{}{3}
		sol.Tests = append(sol.Tests, a)

		a.Input = []interface{}{1, 1}
		a.Output = []interface{}{1}
		sol.Tests = append(sol.Tests, a)

		a.Input = []interface{}{2, 2}
		a.Output = []interface{}{1}
		sol.Tests = append(sol.Tests, a)

		a.Input = []interface{}{1, 2}
		a.Output = []interface{}{0}
		sol.Tests = append(sol.Tests, a)
	*/

	a.Input = []interface{}{2147483647, 2}
	a.Output = []interface{}{1073741823}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0029"] = sol
}
