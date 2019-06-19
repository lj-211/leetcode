package solutions

import (
	"reflect"
)

func fourSumCount(A []int, B []int, C []int, D []int) int {
	one := make(map[int]int)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sum := A[i] + B[j]
			if val, ok := one[sum]; ok {
				one[sum] = val + 1
			} else {
				one[sum] = 1
			}
		}
	}

	ret := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sum := C[i] + D[j]
			search := 0 - sum
			if val, ok := one[search]; ok {
				ret += val
			}
		}
	}

	return ret
}

func init() {
	desc := `
Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.

Example:
Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
	`
	sol := Solution{
		Title:  "4Sum II",
		Desc:   desc,
		Method: reflect.ValueOf(fourSumCount),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0454"] = sol
}
