package main

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

type dpInfo struct {
	Min int
	Max int
}

func absInt(a int) int {
	if a < 0 {
		a *= -1
	}

	return a
}

func dpInternal(node *TreeNode, maxDelta *int) dpInfo {
	if node == nil || maxDelta == nil {
		return dpInfo{
			Min: 0,
			Max: 0,
		}
	}

	left := dpInternal(node.Left, maxDelta)
	right := dpInternal(node.Right, maxDelta)

	minVal := node.Val
	maxVal := node.Val
	if node.Left != nil {
		minVal = minInt(left.Min, minVal)
		maxVal = maxInt(left.Max, maxVal)
	}

	if node.Right != nil {
		minVal = minInt(right.Min, minVal)
		maxVal = maxInt(right.Max, maxVal)
	}

	delta := maxInt(absInt(minVal-node.Val), absInt(maxVal-node.Val))
	if delta > *maxDelta {
		*maxDelta = delta
	}

	return dpInfo{
		Min: minVal,
		Max: maxVal,
	}
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxInterval := 0
	dpInternal(root, &maxInterval)

	return maxInterval
}
