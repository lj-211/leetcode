/*
 * 7. Reverse Integer
 * Given a 32-bit signed integer, reverse digits of an integer.
 *     Example 1:
 *         Input: 123
 *         Output: 321
 *     Example 2:
 *         Input: -123
 *         Output: -321
 *     Example 3:
 *         Input: 120
 *         Output: 21
 */
package solutions

import (
	"fmt"
	"math"
	"time"
)

func reverse(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x /= 10

		if (ret > math.MaxInt32/10) || (ret == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if (ret < math.MinInt32/10) || (ret == math.MinInt32/10 && pop < -8) {
			return 0
		}

		ret = ret*10 + pop
	}
	return ret
}

func reversev2(x int, count *int) int {
	mod := x % 10
	div := x / 10
	if div == 0 {
		return mod
	}

	cnt := reversev2(div, count)
	*count = *count + 1

	val := mod * int(math.Pow10(*count))
	time.Sleep(time.Second)

	return val + cnt
}

func main() {
	input := 1534236469
	//var count int = 0
	//ret := reversev2(input, &count)
	ret := reverse(input)
	fmt.Println(input, " -> ", ret)
	fmt.Println(math.MaxInt32)
}
