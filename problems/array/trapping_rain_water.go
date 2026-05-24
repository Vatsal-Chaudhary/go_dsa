package array

func trap_naive(arr []int) int {
	n := len(arr)

	if n == 0 {
		return 0
	}

	res := 0

	for i := 0; i < n; i++ {
		lMax := arr[i]

		for j := 0; j < i; j++ {
			lMax = max(lMax, arr[j])
		}

		rMax := arr[i]

		for j := i + 1; j < n; j++ {
			rMax = max(rMax, arr[j])
		}

		res += min(lMax, rMax) - arr[i]

	}

	return res
}

func trap_using_leader(arr []int) int {
	n := len(arr)
	if n == 0 || n == 1{
		return 0
	}

	res := 0

	rMax := make([]int, n)
	lMax := make([]int, n)

	lMax[0] = arr[0]
	for i := 1; i < n; i++ {
		lMax[i] = max(lMax[i-1], arr[i])
	}

	rMax[n-1] = arr[n-1]
	for i := n-2; i > 0; i-- {
		rMax[i] = max(rMax[i+1], arr[i])
	}

	for i := 1; i < n-1; i++ {
		res += min(lMax[i], rMax[i]) - arr[i]
	}

	return res
}

func trap(arr []int) int {
	n := len(arr)
	if n == 0 || n == 1 {
		return 0
	}

	left := 0
	right := n - 1
	leftMax := 0
	rightMax := 0
	res := 0

	for left < right {
		if arr[left] < arr[right] {

			if arr[left] >= leftMax {
				leftMax = arr[left]
			} else {
				res += leftMax - arr[left]
			}

			left++
		} else {

			if arr[right] >= rightMax {
				rightMax = arr[right]
			} else {
				res += rightMax - arr[right]
			}

			right--
		}
	}

	return res
}
