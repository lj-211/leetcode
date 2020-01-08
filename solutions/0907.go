package main

import (
	"log"
)

type dpInfo struct {
	Count int
	Min   int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func sumSubarrayMins(A []int) int {
	size := len(A)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return A[0]
	}

	sum := 0
	dp := make([]dpInfo, 0)

	for i := 0; i < size; i++ {
		di := dpInfo{
			Count: 1,
			Min:   A[i],
		}

		for len(dp) > 0 {
			j := len(dp) - 1
			if A[i] <= dp[j].Min {
				di.Count += dp[j].Count
				dp = dp[0:j]
			} else {
				break
			}
		}

		dp = append(dp, di)

		for _, v := range dp {
			sum += (v.Count * v.Min)
		}

		log.Printf("%+v %d\n", dp, sum)
	}

	return sum
}

func main() {
	log.Printf("rst: %d\n", sumSubarrayMins([]int{3, 1, 2, 4}))
}
