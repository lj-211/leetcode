package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func buildTree106(inorder []int, postorder []int) *ds.TreeNode {
	return nil
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
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0106"] = sol
}
