package solutions

import (
	"reflect"
)

type BstNode struct {
	Left     *BstNode
	Right    *BstNode
	LeftSize int
	Cnt      int
	Val      int
}

func insertBst(node *BstNode, val int) int {
	if node == nil {
		return 0
	}

	sum := 0
	for node != nil {
		nval := node.Val
		if val > nval { // right
			sum += (node.LeftSize + node.Cnt)
			if node.Right == nil {
				node.Right = &BstNode{
					Val: val,
					Cnt: 1,
				}
				node = nil
			} else {
				node = node.Right
			}
		} else if val == nval {
			node.Cnt += 1
			sum += node.LeftSize
			node = nil
		} else { // left
			if node.Left == nil {
				node.Left = &BstNode{
					Val: val,
					Cnt: 1,
				}
				node = nil
			} else {
				node.LeftSize += 1
				node = node.Left
			}
		}
	}

	return sum
}

func countSmaller(nums []int) []int {
	size := len(nums)

	if size == 0 {
		return []int{}
	}

	if size < 2 {
		return []int{0}
	}

	ret := make([]int, size)
	root := &BstNode{
		Val: nums[size-1],
		Cnt: 1,
	}

	for i := size - 2; i >= 0; i-- {
		ret[i] = insertBst(root, nums[i])
	}

	return ret
}

func init() {
	desc := `
You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

Example:

Input: [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
	`
	sol := Solution{
		Title:  "Count of Smaller Numbers After Self",
		Desc:   desc,
		Method: reflect.ValueOf(countSmaller),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{5, 2, 6, 1}}
	a.Output = []interface{}{[]int{2, 1, 1, 0}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0315"] = sol
}
