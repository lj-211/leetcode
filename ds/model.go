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
