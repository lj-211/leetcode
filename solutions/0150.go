package solutions

import (
	"reflect"
	"strconv"
)

func calc(one, two, oper string) string {
	ione, _ := strconv.Atoi(one)
	itwo, _ := strconv.Atoi(two)

	rst := 0

	switch {
	case oper == "+":
		rst = ione + itwo
	case oper == "-":
		rst = ione - itwo
	case oper == "*":
		rst = ione * itwo
	case oper == "/":
		rst = ione / itwo
	}

	return strconv.Itoa(rst)
}

func evalRPN(tokens []string) int {
	size := len(tokens)
	if size == 0 {
		return 0
	}

	stack := make([]string, 0)
	operMap := make(map[string]bool)
	operMap["+"] = true
	operMap["-"] = true
	operMap["*"] = true
	operMap["/"] = true

	for i := 0; i < size; i++ {
		v := tokens[i]
		_, ok := operMap[v]
		if !ok {
			stack = append(stack, v)
		} else {
			slen := len(stack)
			cret := calc(stack[slen-2], stack[slen-1], v)
			stack = stack[0 : slen-1]
			stack[slen-2] = cret
		}
	}

	ret, _ := strconv.Atoi(stack[0])

	return ret
}

func init() {
	desc := `
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
	`
	sol := Solution{
		Title:  "逆波兰表达式求值",
		Desc:   desc,
		Method: reflect.ValueOf(evalRPN),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]string{"4", "13", "5", "/", "+"}}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0150"] = sol
}
