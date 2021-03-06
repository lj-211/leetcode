package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func maxProfitIII(prices []int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}

	pre := make([]int, size)
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		pre[i] = ds.MaxInt(pre[i-1], prices[i]-min)
	}

	max := prices[size-1]
	post := make([]int, size)
	for i := size - 2; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		post[i] = ds.MaxInt(post[i+1], max-prices[i])
	}

	max_profit := 0
	for i := 0; i < size; i++ {
		max_profit = ds.MaxInt(max_profit, pre[i]+post[i])
	}

	return max_profit
}

func init() {
	desc := `
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:
Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:
Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
	`
	sol := Solution{
		Title:  "Best Time to Buy and Sell StockIII",
		Desc:   desc,
		Method: reflect.ValueOf(maxProfitIII),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{3, 3, 5, 0, 0, 3, 1, 4}}
	a.Output = []interface{}{6}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 2, 3, 4, 5}}
	a.Output = []interface{}{4}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{7, 6, 4, 3, 1}}
	a.Output = []interface{}{0}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}}
	a.Output = []interface{}{13}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0123"] = sol
}
