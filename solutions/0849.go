package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

func maxDistToClosest(seats []int) int {
	leftp := -1
	rightp := 0

	max := 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			leftp = i
		} else {
			for ; (rightp < len(seats) && seats[rightp] != 1) || rightp < i; rightp++ {
			}
		}

		var max_left, max_right int = 0, 0
		if leftp == -1 {
			max_left = len(seats)
		} else {
			max_left = i - leftp
		}
		if rightp == len(seats) {
			max_right = len(seats)
		} else {
			max_right = rightp - i
		}

		max = ds.MaxInt(max, ds.MinInt(max_left, max_right))
	}

	return max
}

func init() {
	desc := `
In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.
There is at least one empty seat, and at least one person sitting.
Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.
Return that maximum distance to closest person.

Example 1:
	Input: [1,0,0,0,1,0,1]
	Output: 2
	Explanation:
	If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
	If Alex sits in any other open seat, the closest person has distance 1.
	Thus, the maximum distance to the closest person is 2.

Example 2:
	Input: [1,0,0,0]
	Output: 3
	Explanation:
	If Alex sits in the last seat, the closest person is 3 seats away.
	This is the maximum distance possible, so the answer is 3.

Note:
1 <= seats.length <= 20000
seats contains only 0s or 1s, at least one 0, and at least one 1.
	`
	sol := Solution{
		Title:  "Maximize Distance to Closest Person",
		Desc:   desc,
		Method: reflect.ValueOf(maxDistToClosest),
		Tests:  make([]TestCase, 0),
	}
	/*
		a := TestCase{}
		a.Input = []interface{}{[]int{1, 0, 0, 0, 1, 0, 1}}
		a.Output = []interface{}{2}
		sol.Tests = append(sol.Tests, a)
	*/

	a := TestCase{}
	a.Input = []interface{}{[]int{1, 0, 0, 0}}
	a.Output = []interface{}{3}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0849"] = sol
}
