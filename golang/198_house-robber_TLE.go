func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func rob(nums []int) int {

	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return Max(nums[0], nums[1])
	}

	return Max(nums[0]+rob(nums[2:]), nums[1]+rob(nums[3:]))
}