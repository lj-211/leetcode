package solutions

import (
	"reflect"
)

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	s := 1
	e := x
	m := -1

	for s <= e {
		m = (s + e) >> 1
		mpow := m * m
		if mpow < x {
			s = m + 1
		} else if mpow > x {
			e = m - 1
		} else {
			return m
		}
	}

	return e
}

func init() {
	desc := `
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
	`
	sol := Solution{
		Title:  " x 的平方根",
		Desc:   desc,
		Method: reflect.ValueOf(mySqrt),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{8}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0069"] = sol
}
