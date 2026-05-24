package array

func getLargest(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	largest := arr[0]

	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}

	return largest
} 

func secondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	res := arr[0]
	largest := arr[0]

	foundSecond := false

	for _, v := range arr[1:] {
		if v > largest {
			res = largest
			largest = v
			foundSecond = true
		} else if v != largest {
			if !foundSecond || v > res {
			  res = v
			  foundSecond = true
			}
		}
	}

	if !foundSecond {
		return -1
	}

	return res
}
