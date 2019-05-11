package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

type WordDictionary struct {
	Trie *ds.LowcaseAlpabetTrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	trie := &ds.LowcaseAlpabetTrieNode{}
	trie.Cons()
	return WordDictionary{
		Trie: trie,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.Trie.AddWord(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.Trie.SearchWord(word)
}

func TestWordSearch(ws []string, s string) bool {
	w := Constructor()
	for i := 0; i < len(ws); i++ {
		w.AddWord(ws[i])
	}
	return w.Search(s)
}

func init() {
	desc := `
Design a data structure that supports the following two operations:

	void addWord(word)
	bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

Example:
	addWord("bad")
	addWord("dad")
	addWord("mad")
	search("pad") -> false
	search("bad") -> true
	search(".ad") -> true
	search("b..") -> true
Note:
	You may assume that all words are consist of lowercase letters a-z.
	`
	sol := Solution{
		Title:  "Add and Search Word - Data structure design",
		Desc:   desc,
		Method: reflect.ValueOf(TestWordSearch),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]string{"bad", "dad", "mad"}, "pad"}
	a.Output = []interface{}{false}
	sol.Tests = append(sol.Tests, a)

	a = TestCase{}
	a.Input = []interface{}{[]string{"bad", "dad", "mad"}, "b.."}
	a.Output = []interface{}{true}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0211"] = sol
}
