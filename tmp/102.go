/*
 * 102. Binary Tree Level Order Traversal
 * Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * return its level order traversal as:
 * [
 *   [3],
 *   [9,20],
 *   [15,7]
 * ]
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

func recurBt(node *TreeNode, depth int, order_map *[][]int) {
	recurBt(node.Left, depth+1, order_map)
	map_len := len(*order_map)
	if map_len < depth {
		for i := 0; i < depth-map_len; i++ {
			*order_map = append(*order_map, make([]int, 0))
		}
	}
	(*order_map)[depth-1] = append((*order_map)[depth-1], node.Val)
	recurBt(node.Right, depth+1, order_map)
}

func levelOrder(root *TreeNode) [][]int {
	order_map := make([][]int, 0)
	recurBt(root, 1, &order_map)

	return order_map
}

func main() {
	fmt.Println("Binary Tree Level Order Traversal")
}
