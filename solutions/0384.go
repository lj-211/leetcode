package solutions

import (
	//"reflect"
	"math/rand"
)

type Solution0384 struct {
	Original []int
	Array    []int
}

func Constructor0384(nums []int) Solution0384 {
	sol := Solution0384{}
	sol.Original = make([]int, len(nums))
	sol.Array = make([]int, len(nums))
	copy(sol.Original, nums)
	copy(sol.Array, nums)

	return sol
}

/** Resets the array to its original configuration and return it. */
func (this *Solution0384) Reset() []int {
	copy(this.Array, this.Original)
	return this.Array
}

/** Returns a random shuffling of the array. */
func (this *Solution0384) Shuffle() []int {
	size := len(this.Array)
	for i := 0; i < size; i++ {
		sf := rand.Intn(size)
		this.Array[i], this.Array[sf] = this.Array[sf], this.Array[i]
	}
	return this.Array
}

/**
 * Your Solution0384 object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func init() {
	desc := `
打乱一个没有重复元素的数组。

示例:

// 以数字集合 1, 2 和 3 初始化数组。
int[] nums = {1,2,3};
Solution0384 solution = new Solution(nums);

// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
solution.shuffle();

// 重设数组到它的初始状态[1,2,3]。
solution.reset();

// 随机返回数组[1,2,3]打乱后的结果。
solution.shuffle();
	`
	sol := Solution{
		Title: "打乱数组",
		Desc:  desc,
		//Method: reflect.ValueOf(twoSum),
		Tests: make([]TestCase, 0),
	}

	SolutionMap["0384"] = sol
}
