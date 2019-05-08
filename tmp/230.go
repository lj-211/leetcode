/*
 * 230. Kth Smallest Element in a BST
 * Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
 *
 * Example 1:
 * Input: root = [3,1,4,null,2], k = 1
 *    3
 *   / \
 *  1   4
 *   \
 *    2
 * Output: 1
 * Example 2:
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 *        5
 *       / \
 *      3   6
 *     / \
 *    2   4
 *   /
 *  1
 * Output: 3
 * Follow up:
 * What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
 */
package solutions

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrderBst(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	midOrderBst(root.Left, output)
	*output = append(*output, root.Val)
	midOrderBst(root.Right, output)
}

func kthSmallest(root *TreeNode, k int) int {
	output := make([]int, 0)
	midOrderBst(root, &output)

	return output[k-1]
}

func main() {
	tn7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tn8 := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	tn6 := &TreeNode{
		Val:   6,
		Left:  tn7,
		Right: tn8,
	}
	tn5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn3 := &TreeNode{
		Val:   3,
		Left:  tn5,
		Right: tn6,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn4,
		Right: nil,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: tn3,
	}
	/*
		tn3 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		tn2 := &TreeNode{
			Val:   2,
			Left:  tn3,
			Right: nil,
		}
		tn1 := &TreeNode{
			Val:   1,
			Left:  nil,
			Right: tn2,
		}
	*/
	val := kthSmallest(tn1, 4)
	fmt.Println("Kth Smallest Element in a BST: ", val)
}
