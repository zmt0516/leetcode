func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {

	lenn := len(nums)
	switch lenn {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return Max(nums[0], nums[1])
	case 3:
		nums[2] += nums[0]
		return Max(nums[1], nums[2])
	}
	nums[2] += nums[0]
	for i := 3; i < lenn; i++ {
		nums[i] += Max(nums[i-2], nums[i-3])
	}
	return Max(nums[lenn-1], nums[lenn-2])

}