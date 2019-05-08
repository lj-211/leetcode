/*
 * Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 * 	Input: 1->2->4, 1->3->4
 * 	Output: 1->1->2->3->4->4
 */
package solutions

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode = nil
	var cur *ListNode = nil

	for true {
		if l1 == nil && l2 == nil {
			break
		}
		if ret == nil {
			ret = &ListNode{}
			cur = ret
		} else {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
		var v1 int = math.MaxInt32
		var v2 int = math.MaxInt32
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		if v1 <= v2 {
			cur.Val = v1
			l1 = l1.Next
		} else {
			cur.Val = v2
			l2 = l2.Next
		}
	}

	return ret
}

func main() {
	fmt.Println("合并有序数组")
}
