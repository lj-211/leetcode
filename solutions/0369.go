package main

import (
	"fmt"
)

/*
 * 用一个 非空 单链表来表示一个非负整数，然后将这个整数加一。
 *
 * 你可以假设这个整数除了 0 本身，没有任何前导的 0。
 *
 * 这个整数的各个数位按照 高位在链表头部、低位在链表尾部 的顺序排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 */

type Node struct {
	Val int
	Ptr *Node
}

func recurList0369(node *Node) int {
	if node == nil {
		return 1
	}

	step := recurList0369(node.Ptr)
	sum := node.Val + step
	node.Val = sum % 10

	return sum / 10
}

func plusOneList(root *Node) *Node {
	if root == nil {
		return &Node{
			Val: 1,
			Ptr: nil,
		}
	}

	step := recurList0369(root)
	if step > 0 {
		return &Node{
			Val: step,
			Ptr: root,
		}
	}

	return root
}

func main() {
	fmt.Printf("enter main:")

	input := &Node{
		Val: 1,
		Ptr: &Node{
			Val: 2,
			Ptr: &Node{
				Val: 3,
				Ptr: nil,
			},
		},
	}

	input = nil

	input = &Node{
		Val: 9,
		Ptr: &Node{
			Val: 9,
			Ptr: &Node{
				Val: 9,
				Ptr: nil,
			},
		},
	}

	output := plusOneList(input)

	for output != nil {
		fmt.Printf("output: %d\n", output.Val)
		output = output.Ptr
	}

	return
}
