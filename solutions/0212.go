package solutions

import (
//"reflect"
)

func backtracking0212(board [][]byte, mask [][]bool, tree *Trie, i, j int, ilen, jlen int,
	accu string, output *map[string]bool) {
	if tree == nil {
		return
	}
	if mask[i][j] {
		return
	}

	prefix := accu + string(board[i][j])
	ok := tree.StartsWith(prefix)
	if !ok {
		return
	}

	ok = tree.Search(prefix)
	if ok {
		(*output)[prefix] = true
	}

	mask[i][j] = true

	mi := 0
	mj := 0
	// up
	mi = i - 1
	if mi >= 0 && !mask[mi][j] {
		backtracking0212(board, mask, tree, mi, j, ilen, jlen, prefix, output)
	}
	// right
	mj = j + 1
	if mj < jlen && !mask[i][mj] {
		backtracking0212(board, mask, tree, i, mj, ilen, jlen, prefix, output)
	}
	// down
	mi = i + 1
	if mi < ilen && !mask[mi][j] {
		backtracking0212(board, mask, tree, mi, j, ilen, jlen, prefix, output)
	}
	// left
	mj = j - 1
	if mj >= 0 && !mask[i][mj] {
		backtracking0212(board, mask, tree, i, mj, ilen, jlen, prefix, output)
	}

	mask[i][j] = false
}

func findWords(board [][]byte, words []string) []string {
	blen := len(board)
	if blen == 0 {
		return []string{}
	}
	wlen := len(words)
	if wlen == 0 {
		return []string{}
	}
	trie := Constructor0208()
	for i := 0; i < wlen; i++ {
		trie.Insert(words[i])
	}

	mask := make([][]bool, len(board))
	for i := 0; i < len(mask); i++ {
		mask[i] = make([]bool, len(board[0]))
	}

	output := make(map[string]bool)

	for i := 0; i < blen; i++ {
		for j := 0; j < len(board[0]); j++ {
			backtracking0212(board, mask, &trie, i, j, blen, len(board[0]), "", &output)
		}
	}

	oa := make([]string, 0)
	for k, _ := range output {
		oa = append(oa, k)
	}

	return oa
}

func init() {
	desc := `
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

输出: ["eat","oath"]
说明:
你可以假设所有输入都由小写字母 a-z 组成。

提示:

你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
	`
	sol := Solution{
		Title: "单词搜索II",
		Desc:  desc,
		//Method: reflect.ValueOf(findWords),
		Tests: make([]TestCase, 0),
	}

	SolutionMap["0212"] = sol
}
