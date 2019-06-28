package solutions

import (
	"fmt"
	"reflect"
)

const mod int = 1000000007

func countPalindromicSubsequences(S string) int {
	fmt.Println("-----------------")
	slen := len(S)
	if slen == 0 {
		return 0
	}
	if slen == 1 {
		return 1
	}
	dp := make([][]int, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]int, slen)
		dp[i][i] = 1
	}

	// ival != jval
	//	d[i+1][j] + d[i][j-1] - d[i+1][j-1]
	// ival == jval cnt: i+1 - j-1中间ival的数量
	//	cnt == 0
	//		d[i+1][j-1] * 2 + 2
	//	cnt == 1
	//		d[i+1][j-1] * 2 + 1
	//	cnt >= 2
	//		d[i+1][j-1] * 2 - d[r][l]

	for dis := 1; dis < slen; dis++ {
		for i := 0; i < slen-dis; i++ {
			j := i + dis
			ival := S[i]
			jval := S[j]

			if ival != jval {
				dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
				if dp[i][j] < 0 {
					fmt.Println(dp[i+1][j], dp[i][j-1], dp[i+1][j-1])
					fmt.Println("1")
				}
			} else {
				l := i + 1
				r := j - 1
				for l <= r && S[l] != ival {
					l++
				}
				for r >= l && S[r] != jval {
					r--
				}

				if l > r {
					dp[i][j] = dp[i+1][j-1]*2 + 2
				} else if l == r {
					dp[i][j] = dp[i+1][j-1]*2 + 1
				} else {
					dp[i][j] = dp[i+1][j-1]*2 - dp[l+1][r-1]
					if dp[i][j] < 0 {
						fmt.Println("2")
					}
				}
			}

			// dp[i][j]可能为负数 ？ todo: 没搞懂
			// (a - b) % M = (a % M - b % M) + M when a % M - b % M < 0
			if dp[i][j] < 0 {
				dp[i][j] = dp[i][j] + mod
			} else {
				dp[i][j] = dp[i][j] % mod
			}
		}
	}
	return dp[0][slen-1]
}

func init() {
	desc := `
Given a string S, find the number of different non-empty palindromic subsequences in S, and return that number modulo 10^9 + 7.

A subsequence of a string S is obtained by deleting 0 or more characters from S.

A sequence is palindromic if it is equal to the sequence reversed.

Two sequences A_1, A_2, ... and B_1, B_2, ... are different if there is some i for which A_i != B_i.

Example 1:

Input:
S = 'bccb'
Output: 6
Explanation:
The 6 different non-empty palindromic subsequences are 'b', 'c', 'bb', 'cc', 'bcb', 'bccb'.
Note that 'bcb' is counted only once, even though it occurs twice.
Example 2:

Input:
S = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
Output: 104860361
Explanation:
There are 3104860382 different non-empty palindromic subsequences, which is 104860361 modulo 10^9 + 7.
Note:

The length of S will be in the range [1, 1000].
Each character S[i] will be in the set {'a', 'b', 'c', 'd'}.
	`
	sol := Solution{
		Title:  "Count Different Palindromic Subsequences",
		Desc:   desc,
		Method: reflect.ValueOf(countPalindromicSubsequences),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"bccb"}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{"aba"}
	a.Output = []interface{}{4}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{"dbcbaaacdcbabcbddaac"}
	a.Output = []interface{}{356}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{"bddaabdbbccdcdcbbdbddccbaaccabbcacbadbdadbccddccdbdbdbdabdbddcccadddaaddbcbcbabdcaccaacabdbdaccbaacc"}
	a.Output = []interface{}{356}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0730"] = sol
}
