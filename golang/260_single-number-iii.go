func singleNumber(nums []int) []int {
	inc := 0
	for len(nums) > 2 {
		if nums[inc] == nums[inc+1] {
			nums = append(nums[:inc], nums[inc+2:]...)
		} else if nums[inc] > nums[inc+1] {
			nums[inc], nums[inc+1] = nums[inc+1], nums[inc]
		} else {
			inc++
		}
		if inc+1 >= len(nums) {
			inc = 0
		}

	}
	return nums

}