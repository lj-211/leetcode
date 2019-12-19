package main

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return MaxInt(rec1[0], rec2[0]) < MinInt(rec1[2], rec2[2]) &&
		MaxInt(rec1[1], rec2[1]) < MinInt(rec1[3], rec2[3])
}
