package solutions

import (
	"reflect"
)

func maxDistToClosest(seats []int) int {
}

func init() {
	desc := `
	`
	sol := Solution{
		Title:  "Maximize Distance to Closest Person",
		Desc:   desc,
		Method: reflect.ValueOf(maxDistToClosest),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{1, 0, 0, 0, 1, 0, 1}}
	a.Output = []interface{}{2}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0849"] = sol
}
