package solutions

import (
	"fmt"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func wrapIntersectionList(a []int, b []int) []int {
	return ds.ListToArr(getIntersectionNode(ds.MakeListNode(a), ds.MakeListNode(b)))
}

func getIntersectionNode(headA, headB *ds.ListNode) *ds.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	var short, long *ds.ListNode = nil, nil

	step := 0
	for true {
		if a.Next == nil && b.Next == nil {
			break
		}
		if a.Next == nil || b.Next == nil {
			step++
			if a.Next == nil && b.Next != nil {
				short = headA
				long = headB
			}
			if b.Next == nil && a.Next != nil {
				short = headB
				long = headA
			}
		}
		if a.Next != nil {
			a = a.Next
		}
		if b.Next != nil {
			b = b.Next
		}
	}
	if a != b {
		fmt.Println("指针不相等: ", &a, &b)
		return nil
	}

	for step > 0 {
		long = long.Next
		step--
	}

	for long != short {
		long = long.Next
		short = short.Next
	}

	return long
}

func init() {
	desc := `
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:
	`
	sol := Solution{
		Title:  "Intersection of Two Linked Lists",
		Desc:   desc,
		Method: reflect.ValueOf(wrapIntersectionList),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}}
	a.Output = []interface{}{8}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0160"] = sol
}
