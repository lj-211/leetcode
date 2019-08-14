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
func detectCycle(head *ds.ListNode) *ds.ListNode {
	fast := head
	slow := head

	// 第一次相遇，快的从头开始，不相遇则无环
	meet := false
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
		slow = slow.Next

		if fast == slow {
			meet = true
			break
		}
	}
	if !meet {
		return nil
	}

	fast = head
	// 第二次相遇
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func init() {
	desc := `
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
	`
	sol := Solution{
		Title:  "环形链表 II",
		Desc:   desc,
		Method: reflect.ValueOf(detectCycle),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0142"] = sol
}
