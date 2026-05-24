package array

// Window Sliding Technique used here
func maxSumK(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return 0
	}
	curr := 0

	for i := 0; i < k; i++ {
		curr += arr[i]
	}

	res := curr
	for i := k; i < n; i++ {
		curr = curr + arr[i] - arr[i - k]

		res = max(curr, res)
	}

	return res
}
