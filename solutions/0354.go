package solutions

import (
	"fmt"
	"reflect"
	"sort"
)

func envelopLess(a []int, b []int) bool {
	/*
		if a[0] < b[0] {
			return true
		}
		if a[0] == b[0] && a[1] < b[1] {
			return true
		}
	*/
	if a[0] < b[0] && a[1] < b[1] {
		return true
	}

	return false
}

func maxEnvelopes(envelopes [][]int) int {
	size := len(envelopes)
	if size < 2 {
		return size
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			if envelopes[i][1]-envelopes[j][1] > 0 {
				return true
			}
		}

		return false
	})

	fmt.Println(envelopes)

	minList := make([][]int, size)
	minList[0] = envelopes[0]
	maxRight := 0

	for i := 0; i < size; i++ {
		v := envelopes[i]
		l := 0
		r := maxRight
		for l <= r {
			m := (l + r) >> 1
			if envelopLess(minList[m], v) {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		// 这里注意这里可能会导致中间列表minlist不对，
		// 但是不影响，只要没有改变长度就无所谓
		if l > maxRight {
			if v[1] <= minList[maxRight][1] {
				continue
			}

			maxRight = l
		}

		minList[l] = v

		fmt.Println(minList)
		fmt.Println("最右: ", maxRight)
	}

	return maxRight + 1
}

func init() {
	desc := `
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:

输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
	`
	sol := Solution{
		Title:  "俄罗斯套娃信封问题",
		Desc:   desc,
		Method: reflect.ValueOf(maxEnvelopes),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}

	/*
		a.Input = []interface{}{[][]int{[]int{5, 4}, []int{6, 4}, []int{6, 7}, []int{2, 3}}}
		a.Output = []interface{}{3}
		sol.Tests = append(sol.Tests, a)
	*/

	//[[1,3],[3,5],[6,7],[6,8],[8,4],[9,5]]
	/*
		a.Input = []interface{}{[][]int{[]int{1, 3}, []int{3, 5}, []int{6, 7}, []int{6, 8}, []int{8, 4}, []int{9, 5}}}
		a.Output = []interface{}{3}
		sol.Tests = append(sol.Tests, a)
	*/

	// [[2,100],[3,200],[4,300],[5,500],[5,400],[5,250],[6,370],[6,360],[7,380]]
	a.Input = []interface{}{[][]int{[]int{2, 100}, []int{3, 200}, []int{4, 300}, []int{5, 500}, []int{5, 400}, []int{5, 250}, []int{6, 370}, []int{6, 360}, []int{7, 380}}}
	a.Output = []interface{}{5}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0354"] = sol
}
