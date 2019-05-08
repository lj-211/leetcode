/*
 * 6. ZigZag Conversion
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * And then read line by line: "PAHNAPLSIIGYIR"
 * Write the code that will take a string and make this conversion given a number of rows:
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 * Example 2:
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 */
package solutions

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	get_dh := func(line int) int {
		return numRows - line
	}
	get_th := func(line int) int {
		return line - 1
	}

	ret := ""
	for i := 0; i < numRows && i < len(s); i++ {
		step := 0

		dh := get_dh(i + 1)
		th := get_th(i + 1)
		line := string(s[i])
		idx := i
		for idx < len(s) {
			var span, idx_span int = 0, 0
			if (step % 2) == 0 {
				span = dh - 1
				idx_span = 2 * dh
			} else {
				span = th - 1
				idx_span = 2 * th
			}
			if span != -1 {
				if (idx + idx_span) < len(s) {
					line += string(s[idx+idx_span])
					idx = idx + idx_span
				} else {
					break
				}
			}
			step += 1
		}

		ret += line
	}

	return ret
}

func main() {
	//input := "PAYPALISHIRING"
	input := "AB"
	//input := "PAYPALISHIRING"
	input2 := 1
	fmt.Println(convert(input, input2))
}
