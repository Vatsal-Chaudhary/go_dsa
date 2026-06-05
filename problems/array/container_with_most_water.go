package array

import (
	"math"
)

func maxArea(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	l := 0
	r := len(arr) - 1

	res := math.MinInt

	for l < r {
		if arr[l] < arr[r] {
			curr_area := arr[l] * (r - l)
			res = max(curr_area, res)

			l++
		} else {
			curr_area := arr[r] * (r - l)
			res = max(curr_area, res)

			r--
		}
	}

	return res
}
