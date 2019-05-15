package ds

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) String() string {
	ret := ""
	ptr := s
	for ptr != nil {
		ret += strconv.Itoa(ptr.Val) + "->"
		ptr = ptr.Next
	}
	ret += "nil"

	return ret
}

func MakeListNode(a []int) *ListNode {
	var ret *ListNode = nil
	var val *ListNode = nil
	for i := 0; i < len(a); i++ {
		tmp := &ListNode{
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

func ListToArr(node *ListNode) []int {
	ret := make([]int, 0)
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	return ret
}
