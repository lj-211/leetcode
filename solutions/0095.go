package main

import (
	"fmt"
)

func recurTree(start, end int) []*TreeNode {
	if start > end {
		return nil
	}

	if start == end {
		node := &TreeNode{
			Val:   start,
			Left:  nil,
			Right: nil,
		}

		return []*TreeNode{node}
	}

	ret := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		left := recurTree(start, i-1)
		right := recurTree(i+1, end)

		leftLen := len(left)
		rightLen := len(right)

		if leftLen == 0 {
			for jj := 0; jj < rightLen; jj++ {
				ret = append(ret, &TreeNode{
					Val:   i,
					Left:  nil,
					Right: right[jj],
				})
			}
		} else if len(right) == 0 {
			for ii := 0; ii < leftLen; ii++ {
				ret = append(ret, &TreeNode{
					Val:   i,
					Left:  left[ii],
					Right: nil,
				})
			}
		} else {
			for ii := 0; ii < len(left); ii++ {
				for jj := 0; jj < len(right); jj++ {
					ret = append(ret, &TreeNode{
						Val:   i,
						Left:  left[ii],
						Right: right[jj],
					})
				}
			}
		}
	}

	return ret
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return recurTree(1, n)
}
