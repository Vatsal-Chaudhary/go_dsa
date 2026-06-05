package array

func MaxSubArray(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	
	curr_sum := arr[0]
	max_so_far := arr[0]

	for i := 1; i < len(arr); i++ {
		curr_sum = max(curr_sum + arr[i], arr[i])

		max_so_far = max(curr_sum, max_so_far)
	}

	return max_so_far
}
