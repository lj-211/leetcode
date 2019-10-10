package solutions

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0039_1(t *testing.T) {
	in := []int{2, 3, 6, 7}
	target := 7
	eret := [][]int{
		[]int{7},
		[]int{2, 2, 3},
	}
	ret := combinationSum(in, target)
	log.Println(ret)
	assert.Equal(t, true, reflect.DeepEqual(eret, ret), "check ret")
}
