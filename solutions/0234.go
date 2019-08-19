package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func recur234(node *ds.ListNode, ptr **ds.ListNode) bool {
	if node == nil {
		return true
	}

	ret := recur234(node.Next, ptr)
	pval := (*ptr).Val
	*ptr = (*ptr).Next

	if ret && pval == node.Val {
		return true
	}

	return false
}

func isPalindrome234(head *ds.ListNode) bool {
	ptr := head
	return recur234(head, &ptr)
}

func init() {
	desc := `
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
	`
	sol := Solution{
		Title:  "回文链表",
		Desc:   desc,
		Method: reflect.ValueOf(isPalindrome234),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0234"] = sol
}
