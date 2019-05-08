/*
 * 329. Longest Increasing Path in a Matrix
 * Given an integer matrix, find the length of the longest increasing path.
 * From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).
 *
 * Example 1:
 *
 * Input: nums =
 * [
 *   [9,9,4],
 *   [6,6,8],
 *   [2,1,1]
 * ]
 * Output: 4
 * Explanation: The longest increasing path is [1, 2, 6, 9].
 * Example 2:
 *
 * Input: nums =
 * [
 *   [3,4,5],
 *   [3,2,6],
 *   [2,2,1]
 * ]
 * Output: 4
 * Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
 */
package solutions

import (
	"fmt"
)

type NodeInfo struct {
	Ver    int
	Col    int
	Val    int
	Length int

	Left  *NodeInfo
	Right *NodeInfo
	Up    *NodeInfo
	Down  *NodeInfo

	PathBefore *NodeInfo
}

func (this *NodeInfo) HasIncrese() bool {
	return this.Left != nil || this.Right != nil ||
		this.Up != nil || this.Down != nil
}

func VisisNode(node *NodeInfo) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Left.Val > node.Val {
		path := node.Length + 1
		if (node.Left.Length > 0 && path > node.Left.Length) || node.Left.Length == 0 {
			node.Left.Length = path
			node.Left.PathBefore = node
			VisisNode(node.Left)
		}
	}

	if node.Right != nil && node.Right.Val > node.Val {
		path := node.Length + 1
		if (node.Right.Length > 0 && path > node.Right.Length) || node.Right.Length == 0 {
			node.Right.Length = path
			node.Right.PathBefore = node
			VisisNode(node.Right)
		}
	}

	if node.Up != nil && node.Up.Val > node.Val {
		path := node.Length + 1
		if (node.Up.Length > 0 && path > node.Up.Length) || node.Up.Length == 0 {
			node.Up.Length = path
			node.Up.PathBefore = node
			VisisNode(node.Up)
		}
	}

	if node.Down != nil && node.Down.Val > node.Val {
		path := node.Length + 1
		if (node.Down.Length > 0 && path > node.Down.Length) || node.Down.Length == 0 {
			node.Down.Length = path
			node.Down.PathBefore = node
			VisisNode(node.Down)
		}
	}
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	matrix_node := make([][]*NodeInfo, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix_node[i] = make([]*NodeInfo, len(matrix[i]))
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix_node[i][j] = &NodeInfo{
				Ver:        i,
				Col:        j,
				Val:        matrix[i][j],
				Length:     0,
				PathBefore: nil,
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if i != 0 {
				matrix_node[i][j].Up = matrix_node[i-1][j]
			}

			if j != 0 {
				matrix_node[i][j].Left = matrix_node[i][j-1]
			}

			if j != len(matrix[i])-1 {
				matrix_node[i][j].Down = matrix_node[i][j+1]
			}
			if i != len(matrix)-1 {
				matrix_node[i][j].Right = matrix_node[i+1][j]
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			v := matrix_node[i][j]
			VisisNode(v)
		}
	}

	max_node := &NodeInfo{Length: -1}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix_node[i][j].Length > max_node.Length {
				max_node = matrix_node[i][j]
			}
		}
	}

	ret := make([]int, 0)

	for max_node != nil {
		ret = append(ret, max_node.Val)
		max_node = max_node.PathBefore
	}
	fmt.Println(ret)

	return len(ret)
}

func main() {
	fmt.Println("dynamic")
	input := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}

	output := longestIncreasingPath(input)
	fmt.Println("最长增长子串为: ", output)
}
