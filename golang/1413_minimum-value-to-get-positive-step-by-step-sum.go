func minStartValue(nums []int) int {
	sum := 0
	min := nums[0]
	for _, v := range nums {
		sum += v
		if sum < min {
			min = sum
		}
	}
	if min > 0 {
		return 1
	}
	return 1 - min

}