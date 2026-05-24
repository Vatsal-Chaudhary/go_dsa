package array

func moveZeroes(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] != 0 {
					arr[i], arr[j] = arr[j], arr[i]
				}
			} 
		}
	}

	return arr
}

func moveZeroesToEnd(arr []int) []int {
	count := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count], arr[i] = arr[i], arr[count]
			count++
		}
	}

	return arr
}

