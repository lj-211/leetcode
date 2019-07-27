package solutions

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

func coinChange(coins []int, amount int) int {
	noSol := -1
	size := len(coins)
	if size == 0 {
		return noSol
	}
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	if amount < coins[0] {
		return noSol
	}

	dpMap := make(map[int]int)
	dpMap[0] = 0
	for i := 0; i < size; i++ {
		v := coins[i]
		dpMap[v] = 1
	}
	for i := coins[0]; i <= amount; i++ {
		minCnt := math.MaxInt32
		for j := 0; j < size; j++ {
			cv := coins[j]
			dpFind := i - cv
			dp, ok := dpMap[dpFind]
			if ok {
				cnt := dp + 1
				if cnt < minCnt {
					minCnt = cnt
				}
			}
		}
		if minCnt != math.MaxInt32 {
			dpMap[i] = minCnt
		}
	}

	fmt.Println(dpMap)

	cnt, ok := dpMap[amount]
	if ok {
		return cnt
	}

	return noSol
}

func init() {
	desc := `
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
	`
	sol := Solution{
		Title:  "Coin Change",
		Desc:   desc,
		Method: reflect.ValueOf(coinChange),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 5}, 11}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0322"] = sol
}
