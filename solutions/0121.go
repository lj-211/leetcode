package solutions

import (
	"reflect"
)

func maxProfit(prices []int) int {
	profit := make([]int, 0)
	plen := len(prices)
	for i := 1; i < plen; i++ {
		profit = append(profit, prices[i]-prices[i-1])
	}

	return maxSubArray(profit)
}

func init() {
	desc := `
Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
Note that you cannot sell a stock before you buy one.

Example 1:
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
	`
	sol := Solution{
		Title:  "Best Time to Buy and Sell Stock",
		Desc:   desc,
		Method: reflect.ValueOf(maxProfit),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{7, 1, 5, 3, 6, 4}}
	a.Output = []interface{}{5}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0121"] = sol
}
