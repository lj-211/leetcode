package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ds.ListNode, k int) *ds.ListNode {
	if head == nil {
		return head
	}

	size := 0
	node := head
	for node != nil {
		size++
		node = node.Next
	}

	k = k % size

	if k == 0 {
		return head
	}

	first := head
	var second *ds.ListNode = nil
	cnt := 1
	for first.Next != nil {
		cnt++
		first = first.Next

		if cnt == k+1 {
			second = head
		} else if second != nil {
			second = second.Next
		}
	}

	ret := second.Next
	second.Next = nil
	first.Next = head
	return ret
}

func init() {
	desc := `
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
	`
	sol := Solution{
		Title:  "旋转链表",
		Desc:   desc,
		Method: reflect.ValueOf(rotateRight),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0061"] = sol
}
