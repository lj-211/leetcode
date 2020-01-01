/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func robInternal(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	left := robInternal(root.Left)
	right := robInternal(root.Right)

	return []int{
		maxInt(left[0], left[1]) + maxInt(right[0], right[1]),
		root.Val + left[0] + right[0],
	}
}

func rob(root *TreeNode) int {
	ret := robInternal(root)
	return maxInt(ret[0], ret[1])
}
