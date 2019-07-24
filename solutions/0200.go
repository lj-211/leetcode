package solutions

import (
	"fmt"
	"reflect"
)

const (
	Water = 0
	Land  = 1
)

func dfs(grid [][]byte, m, n, i, j int, flag [][]bool) bool {
	v := int(grid[i][j])
	fmt.Println(v, i, j)
	if v == Water || flag[i][j] {
		fmt.Println("ignore: ", i, " - ", j)
		return false
	}

	if v == Land {
		flag[i][j] = true

		if i > 0 && int(grid[i-1][j]) == Land && !flag[i-1][j] {
			dfs(grid, m, n, i-1, j, flag)
		}
		if j < (n-1) && int(grid[i][j+1]) == Land && !flag[i][j+1] {
			dfs(grid, m, n, i, j+1, flag)
		}
		if i < (m-1) && int(grid[i+1][j]) == Land && !flag[i+1][j] {
			dfs(grid, m, n, i+1, j, flag)
		}
		if j > 0 && int(grid[i][j-1]) == Land && !flag[i][j-1] {
			dfs(grid, m, n, i, j-1, flag)
		}
	}

	return true
}

func numIslands(grid [][]byte) int {
	// 1. init
	m := len(grid)
	n := 0
	if m > 0 {
		n = len(grid[0])
	}
	flag := make([][]bool, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
	}

	ret := 0
	// 2. dfs
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rst := dfs(grid, m, n, i, j, flag)
			if rst {
				ret++
			}
		}
	}

	return ret
}

func init() {
	desc := `
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3
	`
	sol := Solution{
		Title:  "nums of island",
		Desc:   desc,
		Method: reflect.ValueOf(numIslands),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[][]uint8{
		[]uint8{1, 1, 1, 1, 0},
		[]uint8{1, 1, 0, 1, 0},
		[]uint8{1, 1, 0, 0, 0},
		[]uint8{0, 0, 0, 0, 0},
	}}
	a.Output = []interface{}{1}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0200"] = sol
}
