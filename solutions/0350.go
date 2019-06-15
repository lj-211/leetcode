package solutions

import (
	"reflect"
	"sort"
)

func intersectionII(nums1 []int, nums2 []int) []int {
	size1 := len(nums1)
	size2 := len(nums2)

	if size1 == 0 || size2 == 0 {
		return []int{}
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	ret := make([]int, 0)
	i, j := 0, 0
	for i < size1 && j < size2 {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ret = append(ret, nums1[1])
			i++
			j++
		}
	}

	return ret
}

func init() {
	desc := `
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.
	`
	sol := Solution{
		Title:  "Intersection of Two Arrays II",
		Desc:   desc,
		Method: reflect.ValueOf(intersection),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 2, 1}, []int{2, 2}}
	a.Output = []interface{}{[]int{2}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0350"] = sol
}
