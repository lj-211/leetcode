package main

import (
	"bufio"
	"fmt"
	"os"
	//"reflect"

	"pokman/bulbasaur/leetcode/solutions"
)

func main() {
	fmt.Println(solutions.SolutionMap)

	f := bufio.NewReader(os.Stdin) //读取输入的内容
	fmt.Print("输入问题序号（例如: 0001）-> ")
	var idx string
	idx, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
	idx = idx[0 : len(idx)-1]
	if sol, ok := solutions.SolutionMap[idx]; ok {
		ret := sol.RunTestCase()
		for i := 0; i < len(ret); i++ {
			fmt.Println(ret[i])
		}
	} else {
		fmt.Println("错误: \n\t1. 问题未收录 \n\t2. 问题未解决")
	}
}
