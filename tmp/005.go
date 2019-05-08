/*
 * 5. Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 * Example 2:
 *
 * Input: "cbbd"
 * Output: "bb"
 */
package solutions

import (
	"fmt"
)

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func longestPalindromeV2(s string) string {
	max_i := 0
	max_j := 0

	map_val := make(map[string]bool)

	str_len := len(s)
	for i := 0; i < str_len; i++ {
		j := i

		for j >= 0 {
			sub_val := fmt.Sprintf("%d-%d", j+1, i-1)
			var sub_bool bool = false
			if val, ok := map_val[sub_val]; ok {
				sub_bool = val
			}
			if s[i] == s[j] && (i-j < 2 && sub_bool) {
				sub_key := fmt.Sprintf("%d-%d", j, i)
				map_val[sub_key] = true
			}

			if i-j > max_i-max_j {
				max_i = i
				max_j = j
			}

			j--
		}
	}

	return s[max_j:max_i]
}

func longestPalindrome(s string) string {
	opt_str := "#"
	for i := 0; i < len(s); i++ {
		opt_str += string(s[i])
		opt_str += "#"
	}

	max_right := 0
	pos := 0
	max_len := 0
	max_str := ""
	var a, b int = 0, 0
	rl := make([]int, len(opt_str))

	for i := 0; i < len(opt_str); i++ {
		if max_right > i {
			rl[i] = Min(rl[2*pos-i], max_right-i)
		} else {
			rl[i] = 1
		}

		for (i-rl[i] >= 0) && (i+rl[i] < len(opt_str)) && opt_str[i-rl[i]] == opt_str[i+rl[i]] {
			rl[i] += 1
		}

		if i+rl[i]-1 > max_right {
			max_right = i + rl[i] - 1
			pos = i
		}

		if rl[i] > max_len {
			max_len = rl[i]
			a = i - rl[i] + 1
			b = 2*rl[i] - 1
			max_str = opt_str[a : a+b]
		}
	}

	ret := ""
	for i := 0; i < len(max_str); i++ {
		if string(max_str[i]) == "#" {
			continue
		}
		ret += string(max_str[i])
	}

	return ret
}

func main() {
	fmt.Println("Longest Palindromic Substring")
	//input := "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"
	input := "aaabaaaa"
	fmt.Println("input: ", input)
	//fmt.Println("最长回文子串: ", longestPalindrome(input))
	fmt.Println("最长回文子串: ", longestPalindromeV2(input))
}
