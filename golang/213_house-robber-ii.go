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
		return Max(Max(nums[0], nums[1]), nums[2])
	case 4:
		return Max(nums[0]+nums[2], nums[1]+nums[3])

	}
	nums2 := make([]int, lenn)
	copy(nums2, nums)
	nums[2] += nums[0]
	for i := 3; i < lenn-1; i++ {
		nums[i] += Max(nums[i-2], nums[i-3])
	}
	max := Max(nums[lenn-2], nums[lenn-3])

	nums2[1] += nums2[lenn-1]
	nums2[2] += Max(nums2[0], nums2[lenn-1])
	for i := 3; i < lenn-2; i++ {
		nums2[i] += Max(nums2[i-2], nums2[i-3])

	}
	max = Max(Max(max, nums2[lenn-4]), nums2[lenn-3])
	return max

}