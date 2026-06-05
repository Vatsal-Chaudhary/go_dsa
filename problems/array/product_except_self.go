package array

func ProductExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}


	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i - 1] * nums[i - 1]
	}

	rightProduct := 1

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= rightProduct

		rightProduct *= nums[i]
	}

	return res
}
