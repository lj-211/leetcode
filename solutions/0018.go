package solutions

import (
	"reflect"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	size := len(nums)

	sort.Ints(nums)

	ret := make([][]int, 0)

	for i := 0; i < size-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < size-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			sum := target - (nums[i] + nums[j])

			s := j + 1
			e := size - 1

			for s < e {
				sval := nums[s]
				eval := nums[e]
				rsum := sval + eval
				if sum == rsum {
					one := make([]int, 0)
					one = append(one, nums[i], nums[j], nums[s], nums[e])
					ret = append(ret, one)
					for s < e {
						s++
						if nums[s] != sval {
							break
						}
					}
					for s < e {
						e--
						if nums[e] != eval {
							break
						}
					}
				} else if sum < rsum {
					e--
				} else {
					s++
				}
			}
		}
	}

	return ret
}

func init() {
	desc := `
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
	`
	sol := Solution{
		Title:  "四数之和",
		Desc:   desc,
		Method: reflect.ValueOf(fourSum),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 0, -1, 0, -2, 2}, 0}
	a.Output = []interface{}{[][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0018"] = sol
}
