package solutions

import (
	"reflect"
)

func byteToInt(b byte) int {
	return int(b) - int('0')
}

func mul(num string, c int, s int) []int {
	// opt: c == 0
	size := len(num)
	ret := make([]int, size+1+s)

	step := 0
	j := s
	for i := size - 1; i >= 0; i-- {
		v := num[i]
		m := byteToInt(v)*c + step
		newVal := m % 10
		step = m / 10
		ret[j] = newVal
		j++
	}
	if step > 0 {
		ret[j] = step
	}

	return ret
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	size2 := len(num2)

	mlist := make([][]int, size2)

	for i := size2 - 1; i >= 0; i-- {
		one := mul(num1, byteToInt(num2[i]), size2-1-i)
		mlist = append(mlist, one)
	}

	// 累加 + 翻转
	ret := ""
	step := 0
	idx := 0
	for {
		sum := 0
		doSum := false

		for i := 0; i < len(mlist); i++ {
			mint := mlist[i]
			mlen := len(mint)
			if idx < mlen {
				doSum = true
				sum += mint[idx]
			}
		}

		if !doSum {
			break
		}

		sum += step
		step = sum / 10
		val := sum % 10
		ret = string(byte(val+int('0'))) + ret

		idx++
	}
	if step > 0 {
		ret = string(byte(step+int('0'))) + ret
	} else {
		if ret[0] == '0' {
			ret = ret[1:]
		}
	}

	return ret
}

func init() {
	desc := `
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
	`
	sol := Solution{
		Title:  "字符串相乘",
		Desc:   desc,
		Method: reflect.ValueOf(multiply),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"123", "456"}
	a.Output = []interface{}{"56088"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0043"] = sol
}

func multiplyV2(num1 string, num2 string) string {
	size1 := len(num1)
	size2 := len(num2)

	if size1 == 0 || size2 == 0 {
		return ""
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	intsList := make([][]int, 0)
	for i := size1 - 1; i >= 0; i-- {
		ival, _ := strconv.Atoi(num1[i : i+1])
		step := 0
		ints := make([]int, size1-1-i)
		for j := size2 - 1; j >= 0; j-- {
			jval, _ := strconv.Atoi(num2[j : j+1])
			mulVal := ival*jval + step
			left := mulVal % 10
			step = mulVal / 10
			ints = append(ints, left)
		}
		if step > 0 {
			ints = append(ints, step)
		}
		intsList = append(intsList, ints)
	}

	sizeList := make([]int, len(intsList))
	for i, v := range intsList {
		sizeList[i] = len(v)
	}

	rstIntList := make([]int, 0)
	col := 0
	step := 0
	for {
		hit := false
		sum := 0
		for i := 0; i < len(intsList); i++ {
			val := intsList[i]
			size := sizeList[i]
			if size <= col {
				continue
			}

			hit = true
			sum += val[col]
		}

		if !hit {
			break
		}
		col++

		sum += step
		cur := sum % 10
		step = sum / 10
		rstIntList = append(rstIntList, cur)
	}
	if step > 0 {
		rstIntList = append(rstIntList, step)
	}

	log.Println(rstIntList)

	ret := ""
	for i := len(rstIntList) - 1; i >= 0; i-- {
		ret += strconv.Itoa(rstIntList[i])
	}

	return ret
}
