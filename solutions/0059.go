package solutions

import (
	"reflect"
)

const (
	Right59 = 0
	Down59  = 1
	Left59  = 2
	Up59    = 3
	Max59   = 4
)

func turnDir(i, j, n, curDir int, mask [][]int) (int, int, int) {
	newi := i
	newj := j
	newDir := curDir

	for {
		dsti := i
		dstj := j

		switch newDir {
		case Left59:
			dsti = i + 1
		case Down59:
			dstj = j + 1
		case Right59:
			dsti = i - 1
		case Up59:
			dstj = j - 1
		}
		if dsti >= 0 && dsti < n && dstj >= 0 && dstj < n && mask[dsti][dstj] == 0 {
			newi = dsti
			newj = dstj
			break
		} else {
			newDir = (newDir + 1) % Max59
			if newDir == curDir {
				break
			}
			continue
		}
	}

	return newi, newj, newDir
}

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	i := 0
	j := 0
	val := 1
	dir := Right59

	for {
		ret[i][j] = val

		newi, newj, newDir := turnDir(i, j, n, dir, ret)
		if newi == i && newj == j {
			break
		}
		i = newi
		j = newj
		dir = newDir
		val++
	}

	return ret
}

func init() {
	desc := `
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
	`
	sol := Solution{
		Title:  "Two Sum",
		Desc:   desc,
		Method: reflect.ValueOf(generateMatrix),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0059"] = sol
}
