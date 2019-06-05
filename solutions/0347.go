package solutions

import (
	"fmt"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func min_cmp(a, b interface{}) bool {
	if a.(Info).Count < b.(Info).Count {
		return true
	}

	return false
}

type Info struct {
	Key   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	fmt.Println("----")
	nmap := make(map[int]int)
	for _, v := range nums {
		if c, ok := nmap[v]; ok {
			nmap[v] = c + 1
		} else {
			nmap[v] = 1
		}
	}
	arr := make([]interface{}, 0)
	for k, v := range nmap {
		arr = append(arr, Info{
			Key:   k,
			Count: v,
		})
	}
	max_heap := ds.BuildHeap(arr, min_cmp)
	fmt.Println("max heap: ", max_heap)

	for len(max_heap) > k {
		heap, _ := ds.PopTop(max_heap, min_cmp)
		max_heap = heap
		//fmt.Println("pop -> ", v.(Info).Key)
	}
	fmt.Println("max heap: ", max_heap)

	ret := make([]int, 0)
	for _, v := range max_heap {
		ret = append(ret, v.(Info).Key)
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
