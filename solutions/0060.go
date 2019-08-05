package solutions

import (
	"fmt"
	"reflect"
	"strconv"
)

func assist(n int) int {
	ret := 1
	for n >= 1 {
		ret *= n
		n--
	}

	return ret
}

func dfs60(nums []int, used []bool, n int, depth int, dst int, selected *[]int) {
	if depth == n {
		return
	}
	step := assist(n - 1 - depth)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if step < dst {
			dst -= step
			continue
		}
		fmt.Println("append: ", nums[i], nums)
		*selected = append(*selected, nums[i])
		used[i] = true

		dfs60(nums, used, n, depth+1, dst, selected)
	}
}

func getPermutation(n int, k int) string {
	ret := ""

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	fmt.Println("-> ", nums)

	used := make([]bool, n)
	selected := make([]int, 0)

	dfs60(nums, used, n, 0, k, &selected)

	fmt.Println(selected)

	for i := 0; i < len(selected); i++ {
		ret += strconv.Itoa(selected[i])
	}

	return ret
}

func init() {
	desc := `
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
	`
	sol := Solution{
		Title:  "第k个排列",
		Desc:   desc,
		Method: reflect.ValueOf(getPermutation),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{3, 5}
	a.Output = []interface{}{"312"}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{4, 9}
	a.Output = []interface{}{"2314"}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{3, 2}
	a.Output = []interface{}{"132"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0060"] = sol
}
