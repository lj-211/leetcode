package solutions

import (
	"fmt"
	"reflect"
)

var SolutionMap map[string]Solution

func init() {
	SolutionMap = make(map[string]Solution)
}

type Solution struct {
	Title  string
	Desc   string
	Method reflect.Value
	Tests  []TestCase
}

type TestCase struct {
	Input  []interface{}
	Output []interface{}
}

func (s Solution) RunTestCase() []string {
	ret := make([]string, 0)
	for i := 0; i < len(s.Tests); i++ {
		td := s.Tests[i]

		ps := make([]reflect.Value, 0)
		for _, v := range td.Input {
			ps = append(ps, reflect.ValueOf(v))
		}

		rst := s.Method.Call(ps)

		output_str := ""
		for j := 0; j < len(rst); j++ {
			var val interface{}
			switch rst[j].Kind() {
			case reflect.String:
				val = rst[j].String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val = rst[j].Int()
			case reflect.Slice:
				l := rst[j].Len()
				tmp := make([]interface{}, 0)
				for idx := 0; idx < l; idx++ {
					tmp = append(tmp, rst[j].Index(i).Interface())
				}
				val = tmp
			default:
				val = "unsupported"
			}
			output_str += fmt.Sprintf("%+v", val)
			if (j + 1) != len(rst) {
				output_str += " "
			}
		}

		expect_str := ""
		for j := 0; j < len(td.Output); j++ {
			expect_str += fmt.Sprintf("%+v", td.Output[j])
			if (j + 1) != len(td.Output) {
				expect_str += " "
			}
		}

		if output_str != expect_str {
			input_str := ""
			for j := 0; j < len(td.Input); j++ {
				input_str += fmt.Sprintf("%+v", td.Input[j])
				if (j + 1) != len(td.Input) {
					input_str += " "
				}
			}
			ret = append(ret,
				fmt.Sprintf("不通过\n\t输入: %s\n\t输出: %s\n\t期望: %s",
					input_str, output_str, expect_str))
		}
	}

	return ret
}
