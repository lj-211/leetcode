package solutions

import (
	"reflect"
)

func reverseString0344(s []byte) {
	size := len(s)
	if size <= 1 {
		return
	}

	begin := 0
	end := size - 1

	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}

func init() {
	desc := `
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

 

示例 1：

输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
	`
	sol := Solution{
		Title:  "反转字符串",
		Desc:   desc,
		Method: reflect.ValueOf(reverseString0344),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]byte{'h', 'e', 'l', 'l', 'o'}}
	a.Output = []interface{}{[]int{'o', 'l', 'l', 'e', 'h'}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0344"] = sol
}
