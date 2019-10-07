package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0035_1(t *testing.T) {
	in := []int{1, 3, 5, 6}
	target := 5

	ret := searchInsert(in, target)
	assert.Equal(t, 2, ret, "return val must be: 2")
}

func Test_0035_2(t *testing.T) {
	in := []int{}
	target := 5

	ret := searchInsert(in, target)
	assert.Equal(t, -1, ret, "must return error idx")
}

func Test_0035_3(t *testing.T) {
	var in []int = nil
	target := 5

	ret := searchInsert(in, target)
	assert.Equal(t, -1, ret, "must return error idx")
}

func Test_0035_4(t *testing.T) {
	in := []int{1, 3, 5, 6}
	target := 2

	ret := searchInsert(in, target)
	assert.Equal(t, 1, ret, "return val must be: 1")
}
