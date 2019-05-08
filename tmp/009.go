/*
 * 9. Palindrome Number
 * Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
 *
 * Example 1:
 * Input: 121
 * Output: true
 * Example 2:
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
 *
 * Example 3:
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 */
package solutions

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revert := 0
	for x > revert {
		revert = revert*10 + x%10
		x = x / 10
	}

	return x == revert || x == (revert/10)
}

func isPalindromeLJ(x int) bool {
	if x < 0 {
		return false
	}
	// 1. 求出位数
	cols := 0
	tmp := x
	for tmp != 0 {
		cols++
		tmp = tmp / 10
	}
	cols--

	ret := false
	for x >= 10 {
		p := int(math.Pow10(cols))
		high := x / p
		low := x % 10

		x = x % p / 10

		cols = cols - 2

		if high != low {
			break
		}
	}

	if x < 10 {
		ret = true
	}

	return ret
}

func main() {
	fmt.Println("Palindrome Number")

	input := 12210
	fmt.Println(input, "是否是回文: ", isPalindrome(input))
}
