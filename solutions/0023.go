package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func cmpListNode(p1 interface{}, p2 interface{}) bool {
	if p1.(*ds.ListNode).Val < p2.(*ds.ListNode).Val {
		return true
	}

	return false
}

func mergeKLists(lists []*ds.ListNode) *ds.ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	dummy := &ds.ListNode{}
	heap_arr := make([]interface{}, size)
	// 1. build heap
	for i := 0; i < size; i++ {
		if lists[i] != nil {
			heap_arr[i] = lists[i]
		}
	}
	heap := ds.BuildHeap(heap_arr, cmpListNode)
	cur := dummy
	for len(heap) > 0 {
		new_heap, min := ds.PopTop(heap, cmpListNode)
		heap = new_heap
		cur.Next = min.(*ds.ListNode)
		cur = cur.Next
		if min.(*ds.ListNode).Next != nil {
			heap = ds.HeapAdd(heap, min.(*ds.ListNode).Next, cmpListNode)
		}
	}
	return nil
}

func init() {
	desc := `
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
	`
	sol := Solution{
		Title:  "Merge k Sorted Lists",
		Desc:   desc,
		Method: reflect.ValueOf(mergeKLists),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}}
	a.Output = []interface{}{[]int{1, 1, 2, 3, 4, 4, 5, 6}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0023"] = sol
}
