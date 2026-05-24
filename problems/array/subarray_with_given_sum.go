package array

func subarraySumExists(arr []int, sum int) bool {
	if len(arr) == 0 {
		return false
	}

	if sum == 0 {
		return false
	}

	s := 0
	curr := 0

	for e := 0; e < len(arr); e++ {
		curr += arr[e]

		for (sum < curr) && (s < e) {
			curr -= arr[s]
			s++
		}

		if curr == sum {
			return true
		}
	}

	return false
}
