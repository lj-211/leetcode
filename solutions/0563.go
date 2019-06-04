package solutions

import (
	"math"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func recurTilt(node *ds.TreeNode) (sum int, tilt int) {
	sum = 0
	tilt = 0
	if node == nil {
		return
	}

	lsum, ltile := recurTilt(node.Left)
	rsum, rtile := recurTilt(node.Right)

	sum = lsum + rsum + node.Val
	tilt = int(math.Abs(float64(lsum-rsum))) + ltile + rtile

	return
}

func findTilt(root *ds.TreeNode) int {
	if root == nil {
		return 0
	}

	_, tilt := recurTilt(root)

	return tilt
}

func init() {
	desc := `
Given a binary tree, return the tilt of the whole tree.
The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.
The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1
Note:

The sum of node values in any subtree won't exceed the range of 32-bit integer.
All the tilt values won't exceed the range of 32-bit integer.
	`
	sol := Solution{
		Title:  "Binary Tree Tilt",
		Desc:   desc,
		Method: reflect.ValueOf(findTilt),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0563"] = sol
}
