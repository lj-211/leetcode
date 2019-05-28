package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func tree(postorder []int, ps int, pe int, inorder []int, is int, ie int) *ds.TreeNode {
	if ps > pe || is > ie {
		return nil
	}
	if ps == pe || is == ie {
		return &ds.TreeNode{
			Val:   postorder[ps],
			Left:  nil,
			Right: nil,
		}
	}
	f := postorder[pe]
	mid_idx := is
	for i := ie; i >= is; i-- {
		if f == inorder[i] {
			mid_idx = i
			break
		}
	}
	left_len := mid_idx - is
	right_len := ie - mid_idx
	left := tree(postorder, ps, ps+left_len-1, inorder, is, mid_idx-1)
	right := tree(postorder, ps+left_len, pe-1, inorder, mid_idx+1, mid_idx+right_len)

	fnode := &ds.TreeNode{
		Val:   f,
		Left:  left,
		Right: right,
	}

	return fnode
}

func buildTree106(inorder []int, postorder []int) *ds.TreeNode {
	plen := len(postorder)
	ilen := len(inorder)
	if postorder == nil || inorder == nil || plen != ilen || plen == 0 {
		return nil
	}

	return tree(postorder, 0, plen-1, inorder, 0, ilen-1)
}

func init() {
	desc := `
	`
	sol := Solution{
		Title:  "Construct Binary Tree from Inorder and Postorder Traversal",
		Desc:   desc,
		Method: reflect.ValueOf(buildTree106),
		Tests:  make([]TestCase, 0),
	}
	/* no test case
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)
	*/

	SolutionMap["0106"] = sol
}
