package solutions

import (
	"math"
	"reflect"
)

/* 贪心
public int jump(int[] A) {
	int jumps = 0, curEnd = 0, curFarthest = 0;
	for (int i = 0; i < A.length - 1; i++) {
		curFarthest = Math.max(curFarthest, i + A[i]);
		if (i == curEnd) {
			jumps++;
			curEnd = curFarthest;
		}
	}
	return jumps;
}
*/

func jump(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return 1
	}

	// 每个idx的最小跳数
	dp := make([]int, size)
	dpCan := make([]bool, size)
	dp[size-1] = 0
	dpCan[size-1] = true

	for i := size - 2; i >= 0; i-- {
		ival := nums[i]
		idx := 1
		minJump := math.MaxInt32
		for ; idx <= ival && (i+idx) <= size-1; idx++ {
			if dpCan[i+idx] && dp[i+idx]+1 < minJump {
				dp[i] = dp[i+idx] + 1
				minJump = dp[i]
				dpCan[i] = true
			}
		}
	}

	return dp[0]
}

func init() {
	desc := `
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:

Input: [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
    Jump 1 step from index 0 to 1, then 3 steps to the last index.
Note:

You can assume that you can always reach the last index.
	`
	sol := Solution{
		Title:  "Jump Game II",
		Desc:   desc,
		Method: reflect.ValueOf(jump),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	a.Input = []interface{}{[]int{2, 3, 1, 1, 4}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0045"] = sol
}
