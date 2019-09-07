package solutions

import (
	"reflect"
)

func reverseSeg(strs []byte, b, e int) {
	for b < e {
		strs[b], strs[e] = strs[e], strs[b]
		b++
		e--
	}
}

func reverseWords0186(strs []byte) {
	size := len(strs)
	reverseSeg(strs, 0, size-1)

	cursor := 0
	for i := 0; i < size; i++ {
		v := strs[i]
		if v == ' ' {
			// cursor - i-1
			reverseSeg(strs, cursor, i-1)
			cursor = i + 1
		}
	}

	if cursor < size {
		reverseSeg(strs, cursor, size-1)
	}
}

func init() {
	desc := `
Given an input string , reverse the string word by word.

Example:

Input:  ['t','h','e',' ','s','k','y',' ','i','s',' ','b','l','u','e']
Output: ['b','l','u','e',' ','i','s',' ','s','k','y',' ','t','h','e']
Note:

A word is defined as a sequence of non-space characters.
The input string does not contain leading or trailing spaces.
The words are always separated by a single space.
Follow up: Could you do it in-place without allocating extra space?
	`
	sol := Solution{
		Title:  "翻转字符串里的单词 II",
		Desc:   desc,
		Method: reflect.ValueOf(reverseWords0186),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0186"] = sol
}
