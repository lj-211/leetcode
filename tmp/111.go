/*
 * 111. Minimum Depth of Binary Tree
 * Given a binary tree, find its minimum depth.
 * The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
 * Note: A leaf is a node with no children.
 *
 * Example:
 * Given binary tree [3,9,20,null,null,15,7],
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * return its minimum depth = 2.
 *
 * Example:
 * 	[1, 2] depth is 2
 */
package solutions

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var min_l, min_r int = 0, 0
	if root.Left != nil {
		min_l = minDepth(root.Left)
	}
	if root.Right != nil {
		min_r = minDepth(root.Right)
	}
	if root.Left != nil && root.Right != nil {
		return 1 + Min(min_l, min_r)
	}
	return 1 + Max(min_l, min_r)
}

func main() {
	fmt.Println("Minimum Depth of Binary Tree")
}
