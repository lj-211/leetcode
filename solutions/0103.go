package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recurTree(node *ds.TreeNode, depth int, output *[][]int) {
	if node == nil {
		return
	}
	if output == nil {
		return
	}
	size := len(*output)
	s := size - 1
	for s <= depth {
		*output = append(*output, make([]int, 0))
	}
	if depth%2 == 0 {
		(*output)[depth] = append((*output)[depth], node.Val)
	} else {
		newArr := make([]int, len((*output)[depth])+1)
		newArr[0] = node.Val
		for i := 0; i < len((*output)[depth]); i++ {
			newArr[i+1] = (*output)[depth][i]
		}
	}
	if node.Left != nil {
		recurTree(node.Left, depth+1, output)
	}
	if node.Right != nil {
		recurTree(node.Right, depth+1, output)
	}
}

func zigzagLevelOrder(root *ds.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	output := make([][]int, 0)

	recurTree(root, 0, &output)

	return output
}

func init() {
	desc := `
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:

[
  [3],
  [20,9],
  [15,7]
]

	`
	sol := Solution{
		Title:  "Binary Tree Zigzag Level Order Traversal",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0103"] = sol
}
