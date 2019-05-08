/*
 * 10. Regular Expression Matching
 * Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
 *
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters like . or *.
 * Example 1:
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * Example 2:
 *
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
 * Example 3:
 *
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 * Example 4:
 *
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
 * Example 5:
 *
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 */
package solutions

import (
	"fmt"
)

var RstMap map[string]bool = make(map[string]bool)

func isMatch(s string, p string) bool {
	if v, ok := RstMap[s+"->"+p]; ok {
		return v
	}
	if p == "" {
		if s == "" {
			return true
		}
		return false
	}
	ret := false
	if true {
		can_extend := false
		p_str := string(p[0])
		if len(p) >= 2 {
			extend := string(p[1])
			if extend == "*" {
				can_extend = true
			}
		}
		first_match := false
		if s != "" && (s[0] == p[0] || p_str == ".") {
			first_match = true
		}

		if first_match {
			if can_extend {
				ret = isMatch(s[1:], p) || isMatch(s, p[2:])
			} else {
				ret = isMatch(s[1:], p[1:])
			}
		} else {
			if can_extend {
				ret = isMatch(s, p[2:])
			} else {
				ret = false
			}
		}
	}

	RstMap[s+"->"+p] = ret

	return ret
}

func isMatch(s string, p string) bool {
	dp_arr := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		tmp := make([]bool, 0)
		for j := 0; j <= len(p); j++ {
			tmp = append(tmp, false)
		}
		dp_arr[i] = tmp
	}
	dp_arr[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			first_match := (i < len(s)) && (s[i] == p[j] || p[j] == '.')
			if (j+1) < len(p) && p[j+1] == '*' {
				dp_arr[i][j] = dp_arr[i][j+2] || (first_match && dp_arr[i+1][j])
			} else {
				dp_arr[i][j] = first_match && dp_arr[i+1][j+1]
			}
		}
	}

	return dp_arr[0][0]
}

func main() {
	//s := "mississippi"
	//p := "mis*is*p*."

	//s := "aa"
	//p := "a*"

	//s := "aab"
	//p := "c*a*b"

	//s := "a"
	//p := "ab*"

	//s := "bbbba"
	//p := ".*a*a"

	//s := "a"
	//p := ".*..a*"

	s := "aaaaaaaaaaaaab"
	p := "a*a*a*a*a*a*a*a*a*a*c"

	ret := isMatch(s, p)
	fmt.Println("是否匹配: ", ret)
	l := len(s)
	fmt.Println("test: ", s[l:])
}
