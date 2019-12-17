package main

import (
	"log"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonPrefix(strs []string, s int, e int) string {
	if s == e {
		return strs[s]
	}
	mid := (s + e) / 2
	left := commonPrefix(strs, s, mid)
	right := commonPrefix(strs, mid+1, e)

	minLen := minInt(len(left), len(right))
	i := 0
	for i < minLen {
		if left[i] != right[i] {
			break
		}
		i++
	}

	return left[0:i]
}

func longestCommonPrefix(strs []string) string {
	size := len(strs)

	if size == 0 {
		return ""
	}

	if size == 1 {
		return strs[0]
	}

	return commonPrefix(strs, 0, size-1)
}

func main() {
	prefix := longestCommonPrefix([]string{"caa", "", "a", "acb"})
	log.Printf("longest prefix: %s\n", prefix)
}
