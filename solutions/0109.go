package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func getMid(node *ds.ListNode) *ds.ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	var pre *ds.ListNode = nil
	var slow *ds.ListNode = node
	var fast *ds.ListNode = node
	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if pre != nil {
		pre.Next = nil
	}

	return slow
}

func sortedListToBST(head *ds.ListNode) *ds.TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &ds.TreeNode{
			Val: head.Val,
		}
	}

	mid := getMid(head)
	right := mid.Next
	mid.Next = nil

	root := &ds.TreeNode{
		Val: mid.Val,
	}

	left := head
	if *left == *mid {
		left = nil
	}

	l := sortedListToBST(left)
	root.Left = l
	r := sortedListToBST(right)
	root.Right = r

	return root
}

func init() {
	desc := `
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
	`
	sol := Solution{
		Title:  "有序链表转换二叉搜索树",
		Desc:   desc,
		Method: reflect.ValueOf(sortedListToBST),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0109"] = sol
}
