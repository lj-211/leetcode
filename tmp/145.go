/*
 * 94. Binary Tree Postorder Traversal
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
 * Output: [3,2,1]
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 */
package solutions

import (
	"fmt"
	"time"
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
	new_right := true
	for true {
		for new_right && node != nil {
			st.Push(node)
			node = node.Left
			continue
		}
		new_right = false

		node = st.Peek()
		if node == nil {
			break
		}

		is_leaf := (node.Left == nil && node.Right == nil)

		if is_leaf {
			val := st.Pop()
			ret = append(ret, val.Val)
			node = st.Peek()
			for node != nil && (node.Right == val || node.Right == nil) { // is leaf & is right node
				val = st.Pop()
				ret = append(ret, val.Val)
				node = st.Peek()
			}
			continue
		} else {
			if node.Right == nil {
				val := st.Pop()
				ret = append(ret, val.Val)
				node = st.Peek()
			} else {
				new_right = true
				fmt.Println("压入右节点: ", node.Right.Val)
				node = node.Right
			}
		}

		time.Sleep(time.Second)
	}

	return ret
}

func main() {
	/*
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
	*/
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

	ret := inorderTraversal(tn1)
	fmt.Println("后序遍历: ", ret)
}
