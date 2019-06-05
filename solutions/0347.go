package solutions

import (
	"fmt"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func min_cmp(a, b int) bool {
	if a < b {
		return true
	}

	return false
}

func topKFrequent(nums []int, k int) []int {
	fmt.Println("----")
	nmap := make(map[int]int)
	rmap := make(map[int]int)
	for _, v := range nums {
		if c, ok := nmap[v]; ok {
			nmap[v] = c + 1
		} else {
			nmap[v] = 1
		}
	}
	arr := make([]int, 0)
	for k, v := range nmap {
		rmap[v] = k
		arr = append(arr, v)
	}
	max_heap := ds.BuildMaxHeap(arr)
	fmt.Println("max heap: ", max_heap)
	ret := make([]int, 0)
	for k > 0 {
		heap, max := ds.PopMaxTop(max_heap)
		max_heap = heap
		ret = append(ret, rmap[max])
		fmt.Println("max heap: ", max_heap)

		k--
	}

	return ret
}

func init() {
	desc := `
Given a non-empty array of integers, return the k most frequent elements.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
	`
	sol := Solution{
		Title:  "Top K Frequent Elements",
		Desc:   desc,
		Method: reflect.ValueOf(topKFrequent),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 1, 1, 2, 2, 3}, 2}
	a.Output = []interface{}{[]int{1, 2}}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{[]int{1, 2}, 2}
	a.Output = []interface{}{[]int{1, 2}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0347"] = sol
}
