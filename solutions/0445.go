package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func addTwoNumbersWrap(l1 []int, l2 []int) []int {
	return ds.ListToArr(addTwoNumbers(ds.MakeListNode(l1), ds.MakeListNode(l2)))
}

func addTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	l1arr := make([]int, 0)
	l2arr := make([]int, 0)
	for l1 != nil {
		l1arr = append(l1arr, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		l2arr = append(l2arr, l2.Val)
		l2 = l2.Next
	}
	l1size := len(l1arr)
	l2size := len(l2arr)

	dummy := &ds.ListNode{}

	i, j := l1size-1, l2size-1
	carry := 0
	for i >= 0 || j >= 0 {
		var ival, jval int
		if i < 0 {
			ival = 0
		} else {
			ival = l1arr[i]
		}
		if j < 0 {
			jval = 0
		} else {
			jval = l2arr[j]
		}

		sum := ival + jval + carry
		v := sum % 10
		carry = sum / 10
		new_node := &ds.ListNode{
			Val:  v,
			Next: dummy.Next,
		}
		dummy.Next = new_node

		i--
		j--
	}
	if carry > 0 {
		dummy.Next = &ds.ListNode{
			Val:  carry,
			Next: dummy.Next,
		}
	}

	return dummy.Next
}

func init() {
	desc := `
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
	`
	sol := Solution{
		Title:  "Add Two Numbers II",
		Desc:   desc,
		Method: reflect.ValueOf(addTwoNumbersWrap),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{7, 2, 4, 3}, []int{5, 6, 4}}
	a.Output = []interface{}{[]int{7, 8, 0, 7}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0445"] = sol
}
