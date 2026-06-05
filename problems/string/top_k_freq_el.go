package string

import (
	"sort"
)

func TopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	freq_slice := [][]int{}

	for num, count := range freq {
		freq_slice = append(freq_slice, []int{num, count})
	}

	sort.Slice(freq_slice, func(i, j int) bool {
		if freq_slice[i][1] == freq_slice[j][1] {
			return freq_slice[i][0] > freq_slice[j][0]
		}

		return freq_slice[i][1] > freq_slice[j][1]
	})

	res := []int{}

	for i := 0; i < k; i++ {
		res = append(res, freq_slice[i][0])
	}

	return res
}
