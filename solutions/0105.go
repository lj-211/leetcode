package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func tree(preorder []int, ps int, pe int, inorder []int, is int, ie int) *ds.TreeNode {
	if ps > pe || is > ie {
		return nil
	}
	if ps == pe || is == ie {
		return &ds.TreeNode{
			Val:   preorder[ps],
			Left:  nil,
			Right: nil,
		}
	}
	f := preorder[ps]
	mid_idx := is
	for i := is; i <= ie; i++ {
		if f == inorder[i] {
			mid_idx = i
			break
		}
	}
	left_len := mid_idx - is
	right_len := ie - mid_idx
	left := tree(preorder, ps+1, ps+left_len, inorder, is, mid_idx-1)
	right := tree(preorder, ps+left_len+1, pe, inorder, mid_idx+1, mid_idx+right_len)

	fnode := &ds.TreeNode{
		Val:   f,
		Left:  left,
		Right: right,
	}

	return fnode
}

func buildTree105(preorder []int, inorder []int) *ds.TreeNode {
	plen := len(preorder)
	ilen := len(inorder)
	if preorder == nil || inorder == nil || plen != ilen || plen == 0 {
		return nil
	}

	return tree(preorder, 0, plen-1, inorder, 0, ilen-1)
}

func init() {
	desc := `
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:
    3
   / \
  9  20
    /  \
   15   7
	`
	sol := Solution{
		Title:  "Construct Binary Tree from Preorder and Inorder Traversal",
		Desc:   desc,
		Method: reflect.ValueOf(buildTree105),
		Tests:  make([]TestCase, 0),
	}
	/* no test case
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)
	*/

	SolutionMap["0105"] = sol
}
