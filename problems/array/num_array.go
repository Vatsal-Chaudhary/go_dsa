package array

type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{prefixSums: []int{}}
	}

	prefix := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	return NumArray{
		prefixSums: prefix,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.prefixSums[right+1] - n.prefixSums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
