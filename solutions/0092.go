package solutions

import (
	"fmt"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func wrapReverseBetween(a []int, m int, n int) []int {
	fmt.Println("调用: ", a, m, n)
	return listToArr(reverseBetween(makeListNode(a), m, n))
}

func reverseBetween(head *ds.ListNode, m int, n int) *ds.ListNode {
	if m == n {
		return head
	}
	fmt.Println("处理: ", head, m, n)
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
				fmt.Println("设置h")
				s = cur
				h.Next = nil
				cur.Next = nil
			} else if i == n {
				t = nxt
				e = cur
				fmt.Println(s)
				fmt.Println(e)
				fmt.Println(t)
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

	fmt.Println(h, s, e, t)
	h.Next = e
	s.Next = t
	fmt.Println(h)

	return dummy.Next
}

func makeListNode(a []int) *ds.ListNode {
	var ret *ds.ListNode = nil
	var val *ds.ListNode = nil
	for i := 0; i < len(a); i++ {
		tmp := &ds.ListNode{
			Val: a[i],
		}
		if i == 0 {
			ret = tmp
		}
		if val == nil {
			val = tmp
		} else {
			val.Next = tmp
			val = val.Next
		}
	}

	return ret
}

func listToArr(node *ds.ListNode) []int {
	ret := make([]int, 0)
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	return ret
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
