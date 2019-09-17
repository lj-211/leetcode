package solutions

import (
	"reflect"
	"strconv"
)

func calculate(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	doCacl := func(one, two, oper string) string {
		ione, _ := strconv.Atoi(one)
		itwo, _ := strconv.Atoi(two)
		if oper == "*" {
			return strconv.Itoa(ione * itwo)
		} else if oper == "/" {
			return strconv.Itoa(ione / itwo)
		}
		return ""
	}

	queue := make([]string, 0)

	digitalStart := 0
	isLastHighPri := false
	if size == 1 {
		ret, _ := strconv.Atoi(s)
		return ret
	}
	for i := 0; i < size; i++ {
		b := s[i]
		switch {
		case b == '+' || b == '-' || b == '*' || b == '/':
			if digitalStart == i {
				queue = append(queue, string(s[i]))
			} else {
				num := s[digitalStart:i]
				if isLastHighPri {
					size := len(queue)
					two := num
					oper := queue[size-1]
					one := queue[size-2]
					result := doCacl(one, two, oper)
					queue = queue[0 : size-1]
					queue[size-2] = result
				} else {
					queue = append(queue, num)
				}
				queue = append(queue, string(s[i]))
			}
			if b == '+' || b == '-' {
				isLastHighPri = false
			} else {
				isLastHighPri = true
			}
			digitalStart = i + 1
		case b == ' ':
			if digitalStart == i {
				digitalStart++
			} else {
				num := s[digitalStart:i]
				if isLastHighPri {
					size := len(queue)
					two := num
					oper := queue[size-1]
					one := queue[size-2]
					result := doCacl(one, two, oper)
					queue = queue[0 : size-1]
					queue[size-2] = result
				} else {
					queue = append(queue, num)
				}
				digitalStart = i + 1
			}
		default:
			if i == size-1 {
				num := s[digitalStart : i+1]
				if isLastHighPri {
					size := len(queue)
					two := num
					oper := queue[size-1]
					one := queue[size-2]
					result := doCacl(one, two, oper)
					queue = queue[0 : size-1]
					queue[size-2] = result
				} else {
					queue = append(queue, num)
				}
			}
		}
	}
	result, _ := strconv.Atoi(queue[0])
	i := 1
	for i < len(queue) {
		oper := queue[i]
		inum, _ := strconv.Atoi(queue[i+1])
		if oper == "+" {
			result += inum
		} else if oper == "-" {
			result -= inum
		}
		i += 2
	}

	return result
}

func init() {
	desc := `
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

示例 1:

输入: "3+2*2"
输出: 7
示例 2:

输入: " 3/2 "
输出: 1
示例 3:

输入: " 3+5 / 2 "
输出: 5
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
	`
	sol := Solution{
		Title:  "基本计算器 II",
		Desc:   desc,
		Method: reflect.ValueOf(calculate),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"3+2*2"}
	a.Output = []interface{}{7}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0227"] = sol
}
