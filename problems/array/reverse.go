package array

func Reverse_array(arr []int) []int {
	low := 0
	high := len(arr) - 1

	for low < high {
		temp := arr[low]
		arr[low] = arr[high]
		arr[high] = temp
	}

	return arr
}
