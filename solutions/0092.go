package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func wrapReverseBetween(a []int, m int, n int) []int {
	return ds.ListToArr(reverseBetween(ds.MakeListNode(a), m, n))
}

func reverseBetween(head *ds.ListNode, m int, n int) *ds.ListNode {
	if m == n {
		return head
	}
	var h *ds.ListNode = nil
	var t *ds.ListNode = nil
	var s *ds.ListNode = nil
	var e *ds.ListNode = nil

	dummy := &ds.ListNode{}
	dummy.Next = head

	i := 1
	pre := dummy
	cur := head
	var nxt *ds.ListNode = cur.Next
	for nxt != nil {
		if i >= m && i <= n {
			if i == m {
				h = pre
				s = cur
				h.Next = nil
				cur.Next = nil
			} else if i == n {
				t = nxt
				e = cur
				break
			}
			new_nxt := nxt.Next
			nxt.Next = cur
			pre = cur
			cur = nxt
			nxt = new_nxt
			if nxt == nil {
				e = cur
			}
		} else {
			pre = cur
			cur = nxt
			nxt = cur.Next
		}

		i++
	}

	h.Next = e
	s.Next = t

	return dummy.Next
}

func init() {
	desc := `
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
	`
	sol := Solution{
		Title:  "Reverse Linked List II",
		Desc:   desc,
		Method: reflect.ValueOf(wrapReverseBetween),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{3, 5}, 1, 2}
	a.Output = []interface{}{[]int{5, 3}}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 5}, 2, 3}
	a.Output = []interface{}{[]int{1, 3, 2, 5}}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 5}, 1, 2}
	a.Output = []interface{}{[]int{2, 1, 3, 5}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0092"] = sol
}
