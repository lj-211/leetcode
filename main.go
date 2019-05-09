package main

import (
	"bufio"
	"fmt"
	"os"

	"pokman/bulbasaur/leetcode/solutions"
)

const (
	CmdTest     = "test"
	CmdShowDesc = "show"
)

func ShowDesc(idx string) {
	if sol, ok := solutions.SolutionMap[idx]; ok {
		fmt.Println(sol.Title)
		fmt.Println(sol.Desc)
	} else {
		fmt.Println("没有对应的题目")
	}
}

func TestCase(idx string) {
	if sol, ok := solutions.SolutionMap[idx]; ok {
		ret := sol.RunTestCase()
		if len(ret) == 0 {
			fmt.Println("测试通过")
		} else {
			for i := 0; i < len(ret); i++ {
				fmt.Println(ret[i])
			}
		}
	} else {
		fmt.Println("错误: \n\t1. 问题未收录 \n\t2. 问题未解决")
	}
}

func main() {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	fmt.Print("输入问题序号（例如: 0001 跑测试用例 0001d 打印问题描述）-> ")
	var idx string
	idx, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
	idx = idx[0 : len(idx)-1]

	var show_desc bool = false
	if len(idx) > 0 && idx[len(idx)-1] == 'd' {
		show_desc = true
		idx = idx[0 : len(idx)-1]
	}

	if show_desc {
		ShowDesc(idx)
	} else {
		TestCase(idx)
	}
}
