package solutions

import (
	"reflect"
)

func reverseString(blist []byte, b, e int) {
	for b < e {
		blist[b], blist[e] = blist[e], blist[b]
		b++
		e--
	}
}

func cutDupBank(blist []byte) []byte {
	validPos := -1

	size := len(blist)

	for i := 0; i < size && validPos < size-1; i++ {
		v := blist[i]
		if v != ' ' {
			validPos++
			blist[validPos] = v
		} else {
			if validPos >= 0 && blist[validPos] != ' ' {
				validPos++
				blist[validPos] = ' '
			}
		}
	}
	end := validPos
	if end < 0 {
		return []byte{}
	}
	if blist[end] != ' ' {
		end++
	}

	val := make([]byte, 0)
	for i := 0; i < end; i++ {
		val = append(val, blist[i])
	}

	return blist[0:end]
}

func reverseWords(s string) string {
	if s == "" {
		return s
	}
	blist := []byte(s)

	blist = cutDupBank(blist)

	// 1. reverse string
	reverseString(blist, 0, len(blist)-1)

	// 2. reverse again
	start := 0
	size := len(blist)
	for i := 0; i < size-1 && start < size; i++ {
		v := blist[i+1]
		if v == ' ' {
			reverseString(blist, start, i)
			start = i + 2
		}
	}

	reverseString(blist, start, len(blist)-1)

	return string(blist)
}

func init() {
	desc := `
给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 

说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
	`
	sol := Solution{
		Title:  "翻转字符串里的单词",
		Desc:   desc,
		Method: reflect.ValueOf(reverseWords),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"hello world"}
	a.Output = []interface{}{"word hello"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0151"] = sol
}
