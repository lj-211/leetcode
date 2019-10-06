package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0016_1(t *testing.T) {
	in := []int{-1, 2, 1, -4}
	target := 1

	ret := threeSumClosest(in, target)
	assert.Equal(t, 2, ret, "return val must be: 2")
}
