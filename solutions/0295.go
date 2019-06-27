package solutions

import (
	"reflect"
)

type MedianFinder struct {
	Datas []int
}

/** initialize your data structure here. */
func MedianConstructor() MedianFinder {
	return MedianFinder{
		Datas: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 找到最后一个小于num的位置
	size := len(this.Datas)
	if size == 0 {
		this.Datas = append(this.Datas, num)
		return
	}
	lo := 0
	hi := size - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		mid_val := this.Datas[mid]
		if mid_val >= num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	ret := make([]int, 0)
	if hi == -1 {
		ret = append(ret, num)
	}
	for i := 0; i < size; i++ {
		ret = append(ret, this.Datas[i])
		if i == hi {
			ret = append(ret, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	size := len(this.Datas)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return float64(this.Datas[0])
	}
	if size%2 == 0 {
		return float64(this.Datas[size/2-1]+this.Datas[size/2]) / 2.0
	} else {
		return float64(this.Datas[size%2])
	}
	return 0.0
}

func init() {
	desc := `
Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.


Example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
	`
	sol := Solution{
		Title:  "Find Median from Data Stream",
		Desc:   desc,
		Method: reflect.ValueOf(twoSum),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0295"] = sol
}
