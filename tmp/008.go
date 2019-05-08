/*
 * 7. String to Integer (atoi)
 * Implement atoi which converts a string to an integer.
 * The function first discards as many whitespace characters as necessary until the first non-whitespace character is found.
 * Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
 * The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
 * If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters,
 * no conversion is performed.
 *
 * Note:
 *	Only the space character ' ' is considered as whitespace character.
 *  Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
 *  If the numerical value is out of the range of representable values, INT_MAX (2^31 − 1) or INT_MIN (−2^31) is returned.
 */
package solutions

import (
	"fmt"
	"math"
	//"time"
)

func myAtoi(str string) int {
	valid_chars := make(map[string]int)
	valid_chars[" "] = 0
	valid_chars["+"] = 0
	valid_chars["-"] = 0
	valid_chars["0"] = 0
	valid_chars["1"] = 0
	valid_chars["2"] = 0
	valid_chars["3"] = 0
	valid_chars["4"] = 0
	valid_chars["5"] = 0
	valid_chars["6"] = 0
	valid_chars["7"] = 0
	valid_chars["8"] = 0
	valid_chars["9"] = 0
	is_valid_char := func(c string, ignore_space bool) bool {
		if _, ok := valid_chars[c]; ok {
			if c == " " && !ignore_space {
				return false
			}
			return true
		}
		return false
	}

	str_len := len(str)
	var is_minus int = 0
	var int_str string
	for i := 0; i < str_len; i++ {
		char_val := string(str[i])
		if is_minus == 0 {
			valid := is_valid_char(char_val, true)
			if !valid {
				return 0
			} else {
				if char_val == "-" {
					is_minus = 1
				} else if char_val == "+" {
					is_minus = 2
				} else if char_val != " " {
					is_minus = 2
					if int_str == "" && char_val == "0" {
					} else {
						int_str += char_val
					}
				}
			}
		} else {
			valid := is_valid_char(char_val, false)
			if !valid {
				break
			} else {
				if char_val == "+" || char_val == "-" {
					if int_str == "" {
						return 0
					} else {
						break
					}
				}
				if int_str == "" && char_val == "0" {
				} else {
					int_str += char_val
				}
			}
		}
	}

	fmt.Println("字符串: ", int_str)

	var ret int = 0
	get_char_int := func(c byte) int {
		return int(c) - int(byte('0'))
	}
	for i := 0; i < len(int_str); i++ {
		ret = ret*10 + get_char_int(int_str[i])
		if is_minus == 2 {
			if ret > math.MaxInt32 {
				ret = math.MaxInt32
				break
			}
		} else {
			if ret > (math.MaxInt32 + 1) {
				ret = math.MinInt32
				break
			}
		}
	}
	if is_minus == 1 && ret != math.MinInt32 {
		ret = -1 * ret
	}

	return ret
}

func main() {
	//input := "  0000000000012345678"
	//input := "+-2"
	//input := "20000000000000000000"
	input := "-5-"
	val := myAtoi(input)
	fmt.Println(input, " -> ", val)
}
