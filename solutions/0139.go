package solutions

import (
	"reflect"
)

func wordBreak(s string, wordDict []string) bool {
	word_count := len(wordDict)
	word_map := make(map[string]bool)
	for i := 0; i < word_count; i++ {
		word_map[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			pre := dp[j]
			if pre == false {
				continue
			}
			left := s[j : i+1]
			_, ok := word_map[left]
			if ok {
				dp[i+1] = true
				break
			}
		}
	}

	if dp[len(s)] {
		return true
	}

	return false
}

func init() {
	desc := `
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
	`
	sol := Solution{
		Title:  "Word Break",
		Desc:   desc,
		Method: reflect.ValueOf(wordBreak),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"leetcode", []string{"leet", "code"}}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{"a", []string{"a"}}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0139"] = sol
}
