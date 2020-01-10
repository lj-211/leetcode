func maxInt(a, b int) int {
    if a > b {
        return a
    }

    return b
}
func longestPalindrome(s string) int {
    sizeStr := len(s)

    if sizeStr == 0 {
        return 0
    }
    if sizeStr == 1 {
        return 1
    }

    mask := make(map[byte]int)
    dp := make([]int, sizeStr)
    dp[0] = 0
    mask[s[0]] = 1
    maxLen := 0

    for i := 1; i < sizeStr; i++ {
        v := s[i]
        cnt, ok := mask[v]
        dp[i] = dp[i-1]
        if ok {
            mask[v] = cnt + 1
            if mask[v] % 2 == 0 {
                dp[i] += 2
                maxLen = maxInt(dp[i], maxLen)
            }
        } else {
            mask[v] = 1
        }
    }

    if maxLen % 2 == 0 && maxLen < sizeStr {
        maxLen += 1
    }

    return maxLen
}
