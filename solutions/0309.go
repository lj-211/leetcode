package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func maxProfitColldown(prices []int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}

	return 0
}

func init() {
	desc := `
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

Example:
Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
	`
	sol := Solution{
		Title:  "Best Time to Buy and Sell Stock with Colldown",
		Desc:   desc,
		Method: reflect.ValueOf(maxProfitColldown),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 2, 3, 0, 2}}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 2, 4}}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{1, 4, 3}}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	a.Input = []interface{}{[]int{2, 1, 2, 1, 0, 1, 2}}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0309"] = sol
}
