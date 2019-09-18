package solutions

import (
	"reflect"
)

func isSubsequence(s string, t string) bool {
	tsize := len(t)
	ssize := len(s)

	if ssize == 0 {
		return true
	}

	scur := 0
	hit := false
	for i := 0; i < tsize && scur < ssize; i++ {
		tv := t[i]
		if tv == s[scur] {
			scur++
			if scur == ssize {
				hit = true
				break
			}
		}
	}

	return hit
}

func init() {
	desc := `
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

示例 1:
s = "abc", t = "ahbgdc"

返回 true.

示例 2:
s = "axc", t = "ahbgdc"

返回 false.

后续挑战 :

如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
	`
	sol := Solution{
		Title:  "判断子序列",
		Desc:   desc,
		Method: reflect.ValueOf(isSubsequence),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"abc", "ahbgdc"}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0392"] = sol
}
