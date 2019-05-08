/*
 * 14. Longest Common Prefix
 * Write a function to find the longest common prefix string amongst an array of strings.
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 * Example 2:
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 * Note:
 * All given inputs are in lowercase letters a-z.
 */
package solutions

import (
	"fmt"
)

import "sort"

/*
 * 解法一: 排序后只比较第一个和最后一个
 * 解法二: 构造单词树，遍历到第一个节点个数大于1的停止
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Sort(sort.StringSlice(strs))

	cmp1 := strs[0]
	cmp2 := strs[len(strs)-1]

	var cm_str string
	for i := 0; i < len(cmp1) && i < len(cmp2); i++ {
		if cmp1[i] != cmp2[i] {
			break
		}
		cm_str += string(cmp1[i])
	}

	return cm_str
}

func main() {
	fmt.Println("最长公共子串")
	//input := []string{"中国", "美国", "美人", "中国人"}
	input := []string{"flower", "flip", "flop"}
	fmt.Println(input)
	ret := longestCommonPrefix(input)
	fmt.Println("最长公共子串是: ", ret)
}
