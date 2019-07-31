package solutions

import (
	"fmt"
	"reflect"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	size := len(nums)
	if size <= 1 {
		return 0
	}

	sort.Ints(nums)

	min := 0
	max := nums[size-1] - nums[0]

	for min < max {
		mid := (min + max) / 2

		cnt := 0   // 计算小于mid的个数
		start := 0 // 当前最小值的游标，如果发生nums[i]-nums[start] > mid，则游标需要右移
		maxDelta := 0

		fmt.Println("------------- mid: ", mid)
		for i := 0; i < size; i++ {
			for start < size && nums[i]-nums[start] > mid {
				start++
			}
			fmt.Println(nums[i], " - ", nums[start])
			delta := nums[i] - nums[start]
			fmt.Println("delta: ", delta)
			if delta > maxDelta {
				maxDelta = delta
			}

			cnt += (i - start)
		}

		fmt.Println("cnt-> ", cnt)
		fmt.Println("start-> ", start)

		if cnt == k {
			return maxDelta
		} else if cnt < k {
			min = mid + 1
		} else {
			max = mid - 1
		}
		fmt.Println("min-max: ", min, max)
	}

	return min
}

func init() {
	desc := `
Given an integer array, return the k-th smallest distance among all the pairs. The distance of a pair (A, B) is defined as the absolute difference between A and B.

Example 1:

Input:
nums = [1,3,1]
k = 1
Output: 0
Explanation:
Here are all the pairs:
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
Then the 1st smallest distance pair is (1,1), and its distance is 0.
Note:

2 <= len(nums) <= 10000.
0 <= nums[i] < 1000000.
1 <= k <= len(nums) * (len(nums) - 1) / 2.
	`
	sol := Solution{
		Title:  "Find K-th Smallest Pair Distance",
		Desc:   desc,
		Method: reflect.ValueOf(smallestDistancePair),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 3, 1}, 1}
	a.Output = []interface{}{0}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 1, 1}, 2}
	a.Output = []interface{}{0}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{62, 100, 4}, 2}
	a.Output = []interface{}{58}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 1, 1}, 2}
	a.Output = []interface{}{0}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0719"] = sol
}
