/*
 * 19. Remove Nth Node From End of List
 * Given a linked list, remove the n-th node from the end of list and return its head.
 *
 * Example:
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes 1->2->3->5.
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 * Could you do this in one pass?
 */
package solutions

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	var second *ListNode = nil
	var pre *ListNode = nil
	cnt := 0
	for true {
		cur = cur.Next
		cnt++
		if cnt >= n {
			if second == nil {
				second = head
			} else {
				pre = second
				second = second.Next
			}

			pre.Next = second.Next
		}
		if cur == nil {
			break
		}
	}

	if second == nil {
		return head.Next
	}

	return head
}

func main() {
	fmt.Println("no test case")
}
