package solutions

import (
	"fmt"
	"reflect"
)

var GRstMap map[string][]string = make(map[string][]string)

func canBreak(s string, wordDict []string) bool {
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
			left := s[j : i+1]
			_, ok := word_map[left]
			if ok && pre {
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

func recurString(s string, dict map[string]bool, max_len int) []string {
	if s == "" {
		return []string{}
	}

	ret := make([]string, 0)

	scur := 0
	ecur := 1
	slen := len(s)

	for ecur < slen {
		if (ecur + 1 - scur) > max_len {
			break
		}
		word := s[scur : ecur+1]
		if _, ok := dict[word]; ok {
			ret = append(ret, word)
			left := s[ecur+1 : slen]

			var out []string
			var ok bool
			if out, ok = GRstMap[left]; !ok {
				out = recurString(s[ecur+1:slen], dict, max_len)
				GRstMap[left] = out
			}
			for _, v := range out {
				ret = append(ret, v)
			}
		}
		ecur++
	}

	return ret
}

func wordBreakTwoV2(s string, wordDict []string) []string {
	if !canBreak(s, wordDict) {
		return []string{}
	}
	word_map := make(map[string]bool)
	word_count := len(wordDict)
	max_len := 0
	for i := 0; i < word_count; i++ {
		wlen := len(wordDict[i])
		if wlen > max_len {
			max_len = wlen
		}
		word_map[wordDict[i]] = true
	}

	return recurString(s, word_map, max_len)
}
func wordBreakTwo(s string, wordDict []string) []string {
	fmt.Println("word break II")
	word_count := len(wordDict)
	word_map := make(map[string]bool)
	for i := 0; i < word_count; i++ {
		word_map[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	dp_str := make([][]string, len(s)+1)
	dp_str[0] = make([]string, 0)

	dp_map := make(map[string][]string)

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			pre := dp[j]
			left := s[j : i+1]
			_, ok := word_map[left]
			if ok && pre {
				dp[i+1] = true
				if len(dp_str[j]) > 0 {
					for idx := 0; idx < len(dp_str[j]); idx++ {
						tmp := dp_str[j][idx] + " " + left
						dp_str[i+1] = append(dp_str[i+1], tmp)
					}
				} else {
					dp_str[i+1] = append(dp_str[i+1], left)
				}
			}
		}
		dp_map[s[0:i+1]] = dp_str[i+1]
	}

	if dp[len(s)] {
		fmt.Println(dp_str)
		return dp_str[len(s)]
	}
	return []string{}
}

func init() {
	desc := `
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]

Example 2:
Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:
Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
	`
	sol := Solution{
		Title:  "Word Break II",
		Desc:   desc,
		Method: reflect.ValueOf(wordBreakTwoV2),
		Tests:  make([]TestCase, 0),
	}
	//a := TestCase{}
	//a.Input = []interface{}{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}}
	//a.Output = []interface{}{[]string{"cat sand dog cats and dog"}"}
	//sol.Tests = append(sol.Tests, a)

	a := TestCase{}
	a.Input = []interface{}{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}
	a.Output = []interface{}{[]string{}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0140"] = sol
}
