package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func minDistance(word1 string, word2 string) int {
	size1 := len(word1)
	size2 := len(word2)

	dp := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		dp[i] = make([]int, size2+1)
		dp[i][0] = i
	}

	for i := 0; i <= size2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = ds.MinInt(dp[i-1][j-1], ds.MinInt(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}

	return dp[size1][size2]
}

func init() {
	desc := `
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:

输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
	`
	sol := Solution{
		Title:  "编辑距离",
		Desc:   desc,
		Method: reflect.ValueOf(minDistance),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"horse", "ros"}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0071"] = sol
}
