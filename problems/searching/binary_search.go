package searching

// [1, 3, 5, 6, 10]
func binarySearch_interative(arr []int, x int) int {
	if len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func binarySearch(arr []int, x int) int {
	return binarySearch_recursive(arr, 0, len(arr) -1, x)
}

func binarySearch_recursive(arr []int, low , high, x int) int {
	if len(arr) == 0 {
		return -1
	}
	if low > high {
		return -1
	}

	mid := (low+high) / 2

	if arr[mid] == x {
		 return mid
	} else if arr[mid] > x {
		return binarySearch_recursive(arr, low, mid - 1, x)
	} else {
		return binarySearch_recursive(arr, mid + 1, high, x)
	}
}
