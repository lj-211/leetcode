package solutions

import (
	"math"
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
func midOrderTree(node *ds.TreeNode, depth int, depthMax *[]int) {
	if node == nil || depthMax == nil {
		return
	}

	size := len(*depthMax)
	left := 0
	if size < depth {
		left = depth - size
	}
	if left > 0 {
		app := make([]int, left)
		for i := 0; i < left; i++ {
			app[i] = math.MinInt32
		}
		*depthMax = append(*depthMax, app...)
	}

	v := node.Val
	if v > (*depthMax)[depth-1] {
		(*depthMax)[depth-1] = v
	}

	midOrderTree(node.Left, depth+1, depthMax)
	midOrderTree(node.Right, depth+1, depthMax)
}

func largestValues(root *ds.TreeNode) []int {
	ret := make([]int, 0)
	midOrderTree(root, 1, &ret)
	return ret
}

func init() {
	desc := `
您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]
	`
	sol := Solution{
		Title:  "在每个树行中找最大值",
		Desc:   desc,
		Method: reflect.ValueOf(largestValues),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0515"] = sol
}
