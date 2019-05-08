/*
 * 11. Container With Most Water
 * Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
 */
package solutions

import (
	"fmt"
)

func IntMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	var max_area, left, right int = 0, 0, len(height) - 1

	for left < right {
		cur_area := (right - left) * IntMin(height[left], height[right])
		max_area = IntMax(cur_area, max_area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max_area
}

func main() {
	fmt.Println("11. Container With Most Water")

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := maxArea(input)
	fmt.Println("最大蓄水量是: ", output)
}
