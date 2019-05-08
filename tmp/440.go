/*
 * 440. K-th Smallest in Lexicographical Order
 * Given integers n and k, find the lexicographically k-th smallest integer in the range from 1 to n.
 *
 * Note: 1 ≤ k ≤ n ≤ 109.
 *
 * Example:
 *
 * Input:
 * n: 13   k: 2
 *
 * Output:
 * 10
 *
 * Explanation:
 * The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.
 */
package solutions

import (
	"fmt"
)

func findKthNumber(n int, k int) int {
	var curr int64 = 1
	k--

	for k > 0 {
		steps := cal_steps(n, curr, curr+1)
		if steps <= k {
			curr += 1
			k = k - steps
		} else {
			k = k - 1
			curr *= 10
		}
	}

	return int(curr)
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func cal_steps(n int, n1 int64, n2 int64) int {
	var steps int64 = 0

	for n1 <= int64(n) {
		steps += min(int64(n+1), n2) - n1
		n1 = n1 * 10
		n2 = n2 * 10
	}

	return int(steps)
}

func main() {
	fmt.Println("K-th Smallest in Lexicographical Order")
	output := findKthNumber(2, 2)
	fmt.Println("output: ", output)
}
