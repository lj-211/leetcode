package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

/* 翻转链表的解法
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	s := head
	f := head

	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	dst := s
	cmp := s
	if f.Next == nil {
		cmp = s.Next
	} else {
		dst = s.Next
		cmp = dst
	}

	var pre *ListNode = nil
	start := head
	for start != dst {
		next := start.Next
		start.Next = pre
		pre = start
		start = next
	}

	cmpLeft := pre
	fmt.Println(cmpLeft.Val, cmp.Val)
	for cmpLeft != nil && cmp != nil {
		if cmpLeft.Val != cmp.Val {
			return false
		}
		cmpLeft = cmpLeft.Next
		cmp = cmp.Next
	}

	return true
}
*/

func recur234(node *ds.ListNode, ptr **ds.ListNode) bool {
	if node == nil {
		return true
	}

	ret := recur234(node.Next, ptr)
	pval := (*ptr).Val
	*ptr = (*ptr).Next

	if ret && pval == node.Val {
		return true
	}

	return false
}

func isPalindrome234(head *ds.ListNode) bool {
	ptr := head
	return recur234(head, &ptr)
}

func init() {
	desc := `
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
	`
	sol := Solution{
		Title:  "回文链表",
		Desc:   desc,
		Method: reflect.ValueOf(isPalindrome234),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0234"] = sol
}
