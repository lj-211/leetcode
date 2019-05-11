package ds

import (
//"fmt"
)

type LowcaseAlpabetTrieNode struct {
	Children []*LowcaseAlpabetTrieNode
	IsWord   bool
	Alpha    byte
}

func (s *LowcaseAlpabetTrieNode) Cons() {
	s.Children = make([]*LowcaseAlpabetTrieNode, 26)
}

func (s *LowcaseAlpabetTrieNode) AddWord(w string) {
	var cursor *LowcaseAlpabetTrieNode = s
	for i := 0; i < len(w) && cursor != nil; i++ {
		alpha := w[i]
		idx := int(alpha) - int('a')
		dst := cursor.Children[idx]
		if dst == nil {
			cursor.Children[idx] = &LowcaseAlpabetTrieNode{
				Alpha:    alpha,
				Children: make([]*LowcaseAlpabetTrieNode, 26),
			}
		}
		cursor = cursor.Children[idx]
	}
	cursor.IsWord = true
}

func (s *LowcaseAlpabetTrieNode) SearchWord(w string) bool {
	var cursor *LowcaseAlpabetTrieNode = s
	find := true
	for i := 0; i < len(w) && cursor != nil; i++ {
		alpha := w[i]
		if alpha == '.' {
			for k := 0; k < len(cursor.Children); k++ {
				if cursor.Children[k] != nil {
					if cursor.Children[k].SearchWord(w[i+1:]) {
						return true
					}
				}
			}
			return false
		} else {
			idx := int(alpha) - int('a')
			cursor = cursor.Children[idx]
		}
	}
	return find && (cursor != nil && cursor.IsWord)
}
