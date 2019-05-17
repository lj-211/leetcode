package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func isUnivalTree(root *ds.TreeNode) bool {
	stack := make([]*ds.TreeNode, 0)
	val := root
	stack = append(stack, val)

	var raw int = root.Val

	for len(stack) > 0 {
		val := stack[len(stack)-1]
		for val.Left != nil {
			stack = append(stack, val.Left)
			val = val.Left
		}
		var last_pop *ds.TreeNode = nil
		for true {
			if len(stack) == 0 {
				break
			}
			val := stack[len(stack)-1]
			if val.Right != nil && val.Right != last_pop {
				stack = append(stack, val.Right)
				break
			} else {
				last_pop = stack[len(stack)-1]
				if last_pop.Val != raw {
					return false
				}
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return true
}

func init() {
	desc := `
A binary tree is univalued if every node in the tree has the same value.
Return true if and only if the given tree is univalued.

Example 1:
Input: [1,1,1,1,1,null,1]
Output: true

Example 2:
Input: [2,2,2,5,2]
Output: false

Note:
The number of nodes in the given tree will be in the range [1, 100].
Each node's value will be an integer in the range [0, 99].
	`
	sol := Solution{
		Title:  "Univalued Binary Tree",
		Desc:   desc,
		Method: reflect.ValueOf(isUnivalTree),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0965"] = sol
}
