package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func wrapReverseList(a []int) []int {
	return ds.ListToArr(reverseList(ds.MakeListNode(a)))
}

func reverseList(head *ds.ListNode) *ds.ListNode {
	if head.Next == nil {
		return head
	}

	/* recursive
	ret := reverseList(head.Next)

	if head.Next.Next == nil {
		head.Next.Next = head
		head.Next = nil
	}
	*/
	var pre *ds.ListNode = nil
	var cur *ds.ListNode = head
	var nxt *ds.ListNode = head.Next
	for nxt != nil {
		new_cur := nxt
		nxt = nxt.Next
		cur.Next = pre
		pre = cur
		cur = new_cur
	}
	cur.Next = pre

	return cur
}

func init() {
	desc := `
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Follow up:
A linked list can be reversed either iteratively or recursively. Could you implement both?
	`
	sol := Solution{
		Title:  "Reverse Linked List",
		Desc:   desc,
		Method: reflect.ValueOf(wrapReverseList),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 4, 5}}
	a.Output = []interface{}{[]int{5, 4, 3, 2, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0206"] = sol
}
