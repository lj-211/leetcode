/*
 * 94. Binary Tree Inorder Traversal
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 * Input: [1,null,2,3]
 *    1
 *     \
 *      2
 *     /
 *    3
 *
 * Output: [1,3,2]
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 */
package solutions

import (
	"fmt"
)

type Stack struct {
	Data []*TreeNode
}

func (s *Stack) Peek() *TreeNode {
	if len(s.Data) == 0 {
		return nil
	}

	return s.Data[len(s.Data)-1]
}
func (s *Stack) Push(input *TreeNode) {
	fmt.Println("压入: ", input.Val)
	s.Data = append(s.Data, input)
}
func (s *Stack) Pop() *TreeNode {
	if len(s.Data) == 0 {
		return nil
	}

	ret := s.Data[len(s.Data)-1]
	s.Data = s.Data[0 : len(s.Data)-1]
	fmt.Println("弹出: ", ret.Val)
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)

	var node *TreeNode = root
	st := Stack{}
	st.Data = make([]*TreeNode, 0)
	for true {
		if node == nil && len(st.Data) == 0 {
			break
		}
		for node != nil {
			st.Push(node)
			node = node.Left
			continue
		}

		ele := st.Pop()
		if ele == nil {
			break
		}
		ret = append(ret, ele.Val)
		node = ele.Right

		/*
			for true {
				if node.Left == ele {
					if node.Right != nil {
						ele = st.Pop()
						ret = append(ret, ele.Val)
						st.Push(node.Right)
						node = node.Right
						break
					}
				}

				ele = st.Pop()
				ret = append(ret, ele.Val)
				node = st.Peek()
				if node == nil {
					break
				}
			}
		*/
	}

	return ret
}

func main() {
	tn7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tn8 := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	tn6 := &TreeNode{
		Val:   6,
		Left:  tn7,
		Right: tn8,
	}
	tn5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn3 := &TreeNode{
		Val:   3,
		Left:  tn5,
		Right: tn6,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn4,
		Right: nil,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: tn3,
	}
	/*
		tn3 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		tn2 := &TreeNode{
			Val:   2,
			Left:  tn3,
			Right: nil,
		}
		tn1 := &TreeNode{
			Val:   1,
			Left:  nil,
			Right: tn2,
		}
	*/

	ret := inorderTraversal(tn1)
	fmt.Println("中序遍历: ", ret)
}
