package solutions

import (
	"fmt"
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

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
		Method: reflect.ValueOf(rotate),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{}
	a.Output = []interface{}{}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0211"] = sol
}
