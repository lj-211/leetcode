package solutions

import (
	"math"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func tmpFunc() {
	/*
		stack := make([]*ds.TreeNode, 0)
		val := root
		stack = append(stack, val)

		for len(stack) > 0 {
			val := stack[len(stack)-1]
			if val.Left != nil {
				for val.Left != nil {
					stack = append(stack, val.left)
					val = val.Left
				}
			} else if val.Right != nil {
				stack = append(stack, val.Right)
			} else {
				stack = stack[0 : len(stack)-2]
			}
		}
	*/
}

func pathRecur(node *ds.TreeNode, max *int) int {
	if node == nil {
		return 0
	}

	left := pathRecur(node.Left, max)
	right := pathRecur(node.Right, max)

	left = ds.MaxInt(0, left)
	right = ds.MaxInt(0, right)

	*max = ds.MaxInt(*max, right+left+node.Val)

	return ds.MaxInt(left, right) + node.Val
}

func maxPathSum(root *ds.TreeNode) int {
	max := math.MinInt32
	pathRecur(root, &max)
	return max
}

func init() {
	desc := `
Given a non-empty binary tree, find the maximum path sum.
For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:
Input: [1,2,3]

       1
      / \
     2   3

Output: 6

Example 2:
Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
	`
	sol := Solution{
		Title:  "Binary Tree Maximum Path Sum",
		Desc:   desc,
		Method: reflect.ValueOf(maxPathSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0124"] = sol
}
