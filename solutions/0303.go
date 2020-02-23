package main

type NumArray struct {
	dpSum []int
}

func Constructor(nums []int) NumArray {
	size := len(nums)
	if size == 0 {
		return NumArray{}
	}
	dpSum := make([]int, size+1)
	dpSum[0] = 0
	for i := 0; i < size; i++ {
		dpSum[i+1] = dpSum[i] + nums[i]
	}
	return NumArray{
		dpSum: dpSum,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dpSum[j+1] - this.dpSum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
