package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	size := len(digits)
	if size == 0 {
		return []int{1}
	}

	step := 0
	digits[size-1]++

	for i := size - 1; i >= 0; i-- {
		sum := digits[i] + step
		step = sum / 10
		digits[i] = sum % 10
	}

	if step > 0 {
		ret := make([]int, size+1)
		ret[0] = step
		return ret
	}

	return digits
}

func main() {
	input := []int{1, 2, 3}
	input = []int{9, 9, 9}
	output := plusOne(input)
	fmt.Printf("output: %+v\n", output)
	return
}
