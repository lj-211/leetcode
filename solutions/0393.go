package solutions

import (
	"reflect"
)

// 1000000 11000000 11100000 11110000
var mask []int = []int{0x80, 0xe0, 0xf0, 0xf8}
var nmask int = 0xc0

func validUtf8(data []int) bool {
	size := len(data)
	if size == 0 {
		return false
	}

	for i := 0; i < size; {
		step := -1
		ival := data[i]
		for m := 0; m < len(mask); m++ {
			mval := mask[m]
			if mval&ival == (mval & (mval - 1)) {
				step = m + 1
				break
			}
		}

		if step == -1 {
			return false
		}

		for j := 1; j < step; j++ {
			idx := i + j
			if idx >= size {
				return false
			}

			if (nmask & data[idx]) != (nmask & (nmask - 1)) {
				return false
			}
		}

		i += step
	}

	return true
}

func init() {
	desc := `
UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：

对于 1 字节的字符，字节的第一位设为0，后面7位为这个符号的unicode码。
对于 n 字节的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为0，后面字节的前两位一律设为10。剩下的没有提及的二进制位，全部为这个符号的unicode码。
这是 UTF-8 编码的工作方式：

   Char. number range  |        UTF-8 octet sequence
      (hexadecimal)    |              (binary)
   --------------------+---------------------------------------------
   0000 0000-0000 007F | 0xxxxxxx
   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
给定一个表示数据的整数数组，返回它是否为有效的 utf-8 编码。

注意:
输入是整数数组。只有每个整数的最低 8 个有效位用来存储数据。这意味着每个整数只表示 1 字节的数据。

示例 1:

data = [197, 130, 1], 表示 8 位的序列: 11000101 10000010 00000001.

返回 true 。
这是有效的 utf-8 编码，为一个2字节字符，跟着一个1字节字符。
	`
	sol := Solution{
		Title:  "UTF-8 编码验证",
		Desc:   desc,
		Method: reflect.ValueOf(validUtf8),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{240, 162, 138, 147, 145}}
	a.Output = []interface{}{false}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0393"] = sol
}
