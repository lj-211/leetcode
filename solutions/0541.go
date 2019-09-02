package solutions

import (
	"reflect"
)

func reverse0541(s []byte, b, e int, n int) {
	if b >= n || e >= n {
		return
	}

	for b < e {
		s[b], s[e] = s[e], s[b]
		b++
		e--
	}
}

func reverseStr0541(s string, k int) string {
	bs := []byte(s)

	size := len(bs)
	cur := 0
	step := 2 * k

	for (cur + step) < size {
		reverse0541(bs, cur, cur+k-1, size)
		cur += step
	}

	left := size - cur
	end := size - 1
	if left > k {
		end = cur + k - 1
	}
	reverse0541(bs, cur, end, size)

	return string(bs)
}

func init() {
	desc := `
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。
	`
	sol := Solution{
		Title:  "反转字符串 II",
		Desc:   desc,
		Method: reflect.ValueOf(reverseStr0541),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"abcdefg", 2}
	a.Output = []interface{}{"bacdfeg"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0541"] = sol
}
