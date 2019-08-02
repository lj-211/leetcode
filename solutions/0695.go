package solutions

import (
	"math"
	"reflect"
)

func dfsIsland(i, j, m, n int, mask [][]bool, grid [][]int, treeNodeCnt *int) {
	if mask[i][j] || grid[i][j] == 0 {
		return
	}

	mask[i][j] = true
	(*treeNodeCnt)++

	if i >= 1 {
		dfsIsland(i-1, j, m, n, mask, grid, treeNodeCnt)
	}
	if i < m-1 {
		dfsIsland(i+1, j, m, n, mask, grid, treeNodeCnt)
	}
	if j >= 1 {
		dfsIsland(i, j-1, m, n, mask, grid, treeNodeCnt)
	}
	if j < n-1 {
		dfsIsland(i, j+1, m, n, mask, grid, treeNodeCnt)
	}
}
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := 0
	if m > 0 {
		n = len(grid[0])
	}

	if m == 0 || n == 0 {
		return 0
	}

	mask := make([][]bool, m)
	for i := 0; i < m; i++ {
		mask[i] = make([]bool, n)
	}

	maxCnt := math.MinInt32

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 0
			dfsIsland(i, j, m, n, mask, grid, &cnt)
			if cnt > maxCnt {
				maxCnt = cnt
			}
		}
	}

	return maxCnt
}

func init() {
	desc := `
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。
	`
	sol := Solution{
		Title:  "岛屿的最大面积",
		Desc:   desc,
		Method: reflect.ValueOf(maxAreaOfIsland),
		Tests:  make([]TestCase, 0),
	}
	//a := TestCase{}

	SolutionMap["0695"] = sol
}
