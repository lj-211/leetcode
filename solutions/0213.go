package main

import (
	"log"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func robInternal(nums []int) int {
	numSize := len(nums)
	if numSize == 0 {
		return 0
	}

	pre := 0
	ppre := 0

	maxVal := 0

	for i := 0; i < numSize; i++ {
		cur := ppre + nums[i]
		val := maxInt(pre, cur)
		maxVal = maxInt(val, maxVal)
		cur = maxVal
		ppre = pre
		pre = cur
	}

	return maxVal
}

func rob(nums []int) int {
	numsSize := len(nums)
	if numsSize == 0 {
		return 0
	}
	if numsSize == 1 {
		return nums[0]
	}

	return maxInt(robInternal(nums[0:numsSize-1]),
		robInternal(nums[1:numsSize]))
}

func main() {
	log.Printf("打家劫舍II:")
	//input := []int{1, 2}
	input := []int{1, 3, 1, 3, 100}
	log.Printf("最大值: %d\n", rob(input))
}
