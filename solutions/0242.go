package solutions

import (
	"reflect"
	"sort"
)

func isAnagram(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	sort.Slice(tb, func(i, j int) bool {
		return tb[i] < tb[j]
	})

	return string(sb) == string(tb)
}

func init() {
	desc := `
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
	`
	sol := Solution{
		Title:  "Two Sum",
		Desc:   desc,
		Method: reflect.ValueOf(isAnagram),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"anagram", "nagaram"}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0242"] = sol
}
