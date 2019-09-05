package solutions

import (
	"reflect"
)

func needCheck(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
		return true
	}

	return false
}

func convert(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		if 'A' > 'a' {
			return byte(int(b) - (int('A') - int('a')))
		} else {
			return byte(int(b) + (int('a') - int('A')))
		}
	}

	return b
}

func isPalindrome0125(s string) bool {
	e := len(s) - 1
	b := 0

	for b < e {
		for b < len(s) && !needCheck(s[b]) {
			b++
		}
		for e >= 0 && !needCheck(s[e]) {
			e--
		}
		if b < e && convert(s[b]) != convert(s[e]) {
			return false
		}
		b++
		e--
	}

	return true
}

func init() {
	desc := `
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
	`
	sol := Solution{
		Title:  "验证回文串",
		Desc:   desc,
		Method: reflect.ValueOf(isPalindrome0125),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"A man, a plan, a canal: Panama"}
	a.Output = []interface{}{}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0125"] = sol
}
