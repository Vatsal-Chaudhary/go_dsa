package array

func canPartitionIntoThree(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	if sum % 3 != 0 {
		return false
	}

	target := sum / 3

	prefixSum := 0
	partitions := 0

	for i := 0; i < len(arr) - 1; i++ {
		prefixSum += arr[i]

		if prefixSum == target {
			prefixSum = 0
			partitions++

			if partitions == 2 {
				return true
			}
		}
	}

	return false
}
