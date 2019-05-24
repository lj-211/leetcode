package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func oddEvenList(head *ds.ListNode) *ds.ListNode {
	if head == nil {
		return nil
	}
	dummy_odd := &ds.ListNode{Val: -1}
	dummy_even := &ds.ListNode{Val: -1}

	i := 0
	cur := head
	odd := dummy_odd
	even := dummy_even
	for cur != nil {
		i++
		if i%2 == 0 {
			even.Next = cur
			even = cur
		} else {
			odd.Next = cur
			odd = cur
		}

		cur = cur.Next
	}

	odd.Next = dummy_even.Next
	even.Next = nil

	return dummy_odd.Next
}

func wrapOddEvenList(a []int) []int {
	return ds.ListToArr(oddEvenList(ds.MakeListNode(a)))
}

func init() {
	desc := `
Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.
You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL

Example 2:
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL

Note:
The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
	`
	sol := Solution{
		Title:  "Odd Even Linked List",
		Desc:   desc,
		Method: reflect.ValueOf(wrapOddEvenList),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 4, 5}}
	a.Output = []interface{}{[]int{1, 3, 5, 2, 4}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0328"] = sol
}
