package searching

func isPair(arr []int, x int) bool {
	if len(arr) < 2 {
		return false
	}

	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i] + arr[j] == x {
			return true
		} else if arr[i] + arr[j] > x {
			j = j - 1
		} else {
			i = i + 1
		}
	}

	return false
}
