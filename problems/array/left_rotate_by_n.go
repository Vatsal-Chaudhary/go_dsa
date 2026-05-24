package array

func rotateLeftOne(arr []int) []int {
	n := len(arr)

	temp := arr[0]

	for i := 1; i < n; i++ {
		arr[i-1] = arr[i]
	}

	arr[n - 1] = temp

	return arr
}

func leftRotateByN_Naive(arr []int, n int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	for range n {
		rotateLeftOne(arr)
	}

	return arr
}

func leftRotateByN(arr []int, n int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	n = n % len(arr)

	temp := make([]int, n)

	for i := 0; i < n; i++ {
		temp[i] = arr[i]	
	}

	for i := n; i < len(arr); i++ {
		arr[i - n] = arr[i]
	}

	for i := 0; i < n; i++ {
		arr[len(arr) - n + i] = temp[i]
	}

	return arr
}
