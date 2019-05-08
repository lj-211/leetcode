/*
 * 179. Largest Number
 * Given a list of non negative integers, arrange them such that they form the largest number.
 *
 * Example 1:
 * Input: [10,2]
 * Output: "210"
 * Example 2:
 *
 * Input: [3,30,34,5,9]
 * Output: "9534330"
 * Note: The result may be very large, so you need to return a string instead of an integer.
 */
package solutions

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	str_arr := make([]string, len(nums))
	for i := 0; i < len(str_arr); i++ {
		str_arr[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(str_arr, func(i, j int) bool {
		stra := str_arr[i] + str_arr[j]
		strb := str_arr[j] + str_arr[i]
		return stra < strb
	})
	fmt.Println("str_arr: ", str_arr)
	if str_arr[len(str_arr)-1] == "0" {
		return "0"
	}

	var ret string
	for i := len(str_arr) - 1; i >= 0; i-- {
		ret += str_arr[i]
	}

	return ret
}

func main() {
	//input := []int{3, 30, 34, 5, 9}
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	ret := largestNumber(input)
	fmt.Println("return val: ", ret)
}
