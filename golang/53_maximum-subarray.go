func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	min := 0
	sum := 0
	res := nums[0]
	for _, v := range nums {
		sum += v
		if sum-min > res {
			res = sum - min
		}
		if min > sum {
			min = sum
		}

	}
	return res
}