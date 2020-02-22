package main

// 2的指数为分界点
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	ret := make([]int, num+1)
	ret[0] = 0
	ret[1] = 1

	next := 2
	cur := 0

	for i := 2; i <= num; i++ {
		if next == i {
			cur = next
			next = next << 1
			ret[i] = 1
		} else {
			ret[i] = ret[cur] + ret[i-cur]
		}
	}

	return ret
}
