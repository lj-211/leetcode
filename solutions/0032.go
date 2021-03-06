package main

import (
	"fmt"
)

// stack
// 最开始要压入的是上个无效字符
func longestValidParenthesesv1(s string) int {
	size := len(s)
	if size <= 1 {
		return 0
	}

	stack := make([]int, 1)
	stack[0] = -1

	max := 0

	for i := 0; i < size; i++ {
		c := s[i]
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = maxInt(max, i-stack[len(stack)-1])
			}
		}
	}

	return max
}

// dp
func longestValidParenthesesv2(s string) int {
	size := len(s)
	if size <= 1 {
		return 0
	}

	dp := make([]int, size+1)
	max := 0

	// '(' = 0 dp[i] = dp[i-1] + 2
	// '))'
	//		'))' dp[i-dp[i-1]-1] == '(' dp[i-dp[i-1]-2] + 2
	//		''

	for i := 1; i < size; i++ {
		dpIdx := i + 1
		v := s[i-1]
		if s[i] == '(' {
			dp[dpIdx] = 0
		} else {
			if v == '(' {
				dp[dpIdx] = dp[dpIdx-2] + 2
			} else if dpIdx >= (dp[dpIdx-1]+2) && s[dpIdx-dp[dpIdx-1]-2] == '(' {
				dp[dpIdx] = dp[dpIdx-1] + dp[dpIdx-dp[dpIdx-1]-2] + 2
			} else {
				dp[dpIdx] = 0
			}

			if dp[dpIdx] > max {
				max = dp[dpIdx]
			}
		}
	}

	return max
}

// 利用括号特性解法 关键点: 合适left&right清0
// ")))))()()(((((("
func longestValidParenthesesv3(s string) int {
	maxLeft := 0
	maxRight := 0

	size := len(s)
	max := 0
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			maxLeft++
		} else {
			maxRight++
		}

		if maxLeft == maxRight {
			max = maxInt(maxLeft*2, max)
		} else if maxLeft < maxRight {
			maxLeft = 0
			maxRight = 0
		}
	}

	maxLeft = 0
	maxRight = 0
	for i := size - 1; i >= 0; i-- {
		if s[i] == ')' {
			maxRight++
		} else {
			maxLeft++
		}

		if maxLeft == maxRight {
			max = maxInt(maxLeft*2, max)
		} else if maxLeft > maxRight {
			maxLeft = 0
			maxRight = 0
		}
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Printf("32: 最长有效括号\n")

	input := "()))()"
	input = "()()"
	input = "()(()())"
	output := longestValidParenthesesv3(input)

	fmt.Printf("intput: %s\noutput: %d \n", input, output)
}
