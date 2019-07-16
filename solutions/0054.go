package solutions

import (
	"fmt"
	"reflect"
)

const (
	Right = iota
	Down
	Left
	Up
)

func move(i, j int, dir *int, r, d, l, u *int) (int, int) {
	reti, retj := 0, 0

	switch *dir {
	case Right:
		nextj := j + 1
		if nextj == *r {
			*r = j
			*dir = Down
			reti = i + 1
			retj = j
		} else {
			reti = i
			retj = j + 1
		}
	case Down:
		nexti := i + 1
		if nexti == *d {
			*d = i
			*dir = Left
			reti = i
			retj = j - 1
		} else {
			reti = i + 1
			retj = j
		}
	case Left:
		nextj := j - 1
		if nextj == *l {
			*l = j
			*dir = Up
			reti = i - 1
			retj = j
		} else {
			reti = i
			retj = j - 1
		}
	case Up:
		nexti := i - 1
		if nexti == *u {
			*u = i
			*dir = Right
			reti = i
			retj = j + 1
		} else {
			reti = i - 1
			retj = j
		}
	}

	if (reti == *u || reti == *d) && (retj == *r || retj == *l) {
		reti = -1
		retj = -1
	}

	return reti, retj
}

func spiralOrder(matrix [][]int) []int {
	sizeX := len(matrix)
	if sizeX == 0 {
		return []int{}
	}
	sizeY := len(matrix[0])
	if sizeY == 0 {
		return []int{}
	}
	// dir
	r := sizeY
	d := sizeX
	l := -1
	u := 0
	i, j := 0, 0
	dir := Right
	ret := make([]int, 1)
	ret[0] = matrix[0][0]
	for {
		i, j = move(i, j, &dir, &r, &d, &l, &u)
		if i >= 0 && i < sizeX && j >= 0 && j < sizeY {
			fmt.Println("走一步: ", i, j, " 值为: ", matrix[i][j], " 方向: ", dir)
		} else {
			fmt.Println("异常i j: ", i, j, " 方向: ", dir)
		}
		if i == -1 && j == -1 {
			break
		} else {
			ret = append(ret, matrix[i][j])
		}
	}
	return ret
}

func init() {
	desc := `
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
	`
	sol := Solution{
		Title:  "Spiral Matrix",
		Desc:   desc,
		Method: reflect.ValueOf(spiralOrder),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	/*
		a.Input = []interface{}{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}}
		a.Output = []interface{}{[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}}
		sol.Tests = append(sol.Tests, a)
	*/

	/*
		a.Input = []interface{}{[][]int{[]int{1, 2, 3}}}
		a.Output = []interface{}{[]int{1, 2, 3}}
		sol.Tests = append(sol.Tests, a)
	*/

	a.Input = []interface{}{[][]int{[]int{1}, []int{2}, []int{3}}}
	a.Output = []interface{}{[]int{1, 2, 3}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0054"] = sol
}
