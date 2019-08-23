package solutions

import (
	"reflect"
	"strconv"
)

func validMask(s string) bool {
	size := len(s)
	if size > 1 && s[0] == '0' {
		return false
	}
	if size > 3 {
		return false
	}

	val, _ := strconv.Atoi(s)
	if val > 255 {
		return false
	}

	return true
}

func backtracking0093(s string, depth int) ([]string, bool) {
	nilReturn := []string{}
	if s == "" && depth <= 3 {
		return nilReturn, false
	}

	if depth > 3 {
		return nilReturn, false
	}
	size := len(s)
	if size < (4 - depth) {
		return nilReturn, false
	}

	if depth == 3 {
		if size > 3 {
			return nilReturn, false
		} else {
			if validMask(s) {
				return []string{s}, true
			}
			return nilReturn, false
		}
	}

	ret := make([]string, 0)
	for i := 0; i < size && i < 3; i++ {
		val := s[0 : i+1]
		if validMask(val) {
			list, ok := backtracking0093(s[i+1:], depth+1)
			if ok {
				for x := 0; x < len(list); x++ {
					ret = append(ret, val+"."+list[x])
				}
			}
		}
	}
	if len(ret) == 0 {
		return nilReturn, false
	}

	return ret, true
}

func restoreIpAddresses(s string) []string {
	ret, _ := backtracking0093(s, 0)
	return ret
}

func init() {
	desc := `
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
	`
	sol := Solution{
		Title:  "复原IP地址",
		Desc:   desc,
		Method: reflect.ValueOf(restoreIpAddresses),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"0000"}
	a.Output = []interface{}{[]string{"0.0.0.0"}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0093"] = sol
}
