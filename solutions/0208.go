package solutions

import (
//"reflect"
)

const AlphaCnt int = 26

type TrieElem struct {
	Alpha byte
	Word  bool
	Leafs []*TrieElem
}
type Trie struct {
	Root *TrieElem
}

/** Initialize your data structure here. */
func Constructor0208() Trie {
	return Trie{
		Root: &TrieElem{
			Leafs: make([]*TrieElem, AlphaCnt),
		},
	}
}

func getAlphaIdx(b byte) int {
	return int(b) - int('a')
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	size := len(word)

	node := this.Root

	for i := 0; i < size; i++ {
		b := word[i]
		idx := getAlphaIdx(b)
		next := node.Leafs[idx]
		if next == nil {
			node.Leafs[idx] = &TrieElem{
				Alpha: b,
				Leafs: make([]*TrieElem, AlphaCnt),
			}

			next = node.Leafs[idx]
		}

		node = next
	}
	node.Word = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	size := len(word)
	hit := true
	node := this.Root

	for i := 0; i < size; i++ {
		b := word[i]
		idx := getAlphaIdx(b)

		node = node.Leafs[idx]
		if node == nil {
			hit = false
			break
		}
	}

	return hit && node.Word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	size := len(prefix)
	hit := true
	node := this.Root

	for i := 0; i < size; i++ {
		b := prefix[i]
		idx := getAlphaIdx(b)

		node = node.Leafs[idx]
		if node == nil {
			hit = false
			break
		}
	}

	return hit
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func init() {
	desc := `
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
	`
	sol := Solution{
		Title: "实现 Trie (前缀树)",
		Desc:  desc,
		//Method: reflect.ValueOf(twoSum),
		Tests: make([]TestCase, 0),
	}

	SolutionMap["0208"] = sol
}
