package array

func equilibrium(arr []int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}

	rs := 0
	for i := 0; i < n; i++ {
		rs += arr[i]
	}

	ls := 0
	for i := 0; i < n; i++ {
		rs -= arr[i]

		if ls == rs {
			return true
		}

		ls += arr[i]
	}

	return false
}
