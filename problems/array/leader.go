package array

func leadersInArray_Naive(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	res := []int{}

	for i := 0; i < len(arr); i++ {
		flag := false

		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				flag = true
				break
			}
		}

		if flag == false {
			res = append(res, arr[i])
		}
	}

	return res
}

func leadersInArray(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	res := []int{}

	curr_leader := arr[n - 1]

	res = append(res, curr_leader)

	for i := n - 2; i >= 0; i-- {
		if curr_leader < arr[i] {
			curr_leader = arr[i]
			res = append(res, curr_leader)
		}
	}

	return res
}
