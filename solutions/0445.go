package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func addTwoNumbersWrap(l1 []int, l2 []int) []int {
	return []int{}
}

func addTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	return nil
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
