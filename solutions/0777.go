package solutions

import (
	"log"
	"reflect"
)

func canTransform(start string, end string) bool {
	log.Println("in -> ", start, " to -> ", end)

	size := len(start)
	cntl := 0
	cntr := 0

	misMatch := true

	for i := 0; i < size; i++ {
		sv := start[i]
		ev := end[i]

		if sv == 'R' {
			cntr++
		}
		if ev == 'L' {
			cntl++
		}

		if ev == 'R' {
			cntr--
		}
		if sv == 'L' {
			cntl--
		}

		if cntl < 0 || cntr < 0 || cntl*cntr != 0 {
			misMatch = false
			break
		}
	}

	return misMatch && cntr == 0 && cntl == 0
}

func init() {
	desc := `
在一个由 'L' , 'R' 和 'X' 三个字符组成的字符串（例如"RXXLRXRXL"）中进行移动操作。一次移动操作指用一个"LX"替换一个"XL"，或者用一个"XR"替换一个"RX"。现给定起始字符串start和结束字符串end，请编写代码，当且仅当存在一系列移动操作使得start可以转换成end时， 返回True。

示例 :

输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
输出: True
解释:
我们可以通过以下几步将start转换成end:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX
注意:

1 <= len(start) = len(end) <= 10000。
start和end中的字符串仅限于'L', 'R'和'X'。
	`
	sol := Solution{
		Title:  "在LR字符串中交换相邻字符",
		Desc:   desc,
		Method: reflect.ValueOf(canTransform),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"RXXLRXRXL"}
	a.Output = []interface{}{"XRLXXRRLX"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0777"] = sol
}
