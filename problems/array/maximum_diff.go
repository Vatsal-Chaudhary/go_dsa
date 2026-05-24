package array

// Time - O(n * n)
func maxDifference_naive(arr []int) int {
	n := len(arr)

	if n < 2 {
		return -1
	}

	res := arr[1] - arr[0]

	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			res = max(res, arr[j] - arr[i])
		}
	}

	return res
}

func maxDifference(arr []int) int {
	n := len(arr)

	if n < 2 {
		return -1
	}

	res := arr[1] - arr[0]
	minVal := arr[0]

	for i := 1; i < n; i++ {
		res = max(res, arr[i] - minVal)

		minVal = min(minVal, arr[i])
	}

	return res
}
