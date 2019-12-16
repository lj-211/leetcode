package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Llongfile|log.Ldate)

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	size := len(s)

	if size <= 1 {
		return size
	}

	// key: 字符 value: 最新的位置后的一个位置
	posList := make([]int, 256)
	maxLen := 0
	left := 0
	for i := 0; i < size; i++ {
		v := int(s[i])
		left = MaxInt(posList[v], left)
		maxLen = MaxInt(maxLen, i-left+1)
		posList[v] = i + 1
	}

	return maxLen
}

func main() {
	maxLen := lengthOfLongestSubstring("tmmzxya")
	logger.Printf("最长非重复子串: %d\n", maxLen)
}
