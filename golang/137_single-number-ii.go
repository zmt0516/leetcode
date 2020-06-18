func singleNumber(nums []int) int {

	inc := 0
	for len(nums) > 1 {
		if nums[inc] == nums[inc+1] && nums[inc+1] == nums[inc+2] {
			nums = append(nums[:inc], nums[inc+3:]...)
		} else {
			if nums[inc] > nums[inc+1] {
				nums[inc], nums[inc+1] = nums[inc+1], nums[inc]
			}
			if nums[inc+1] > nums[inc+2] {
				nums[inc+1], nums[inc+2] = nums[inc+2], nums[inc+1]
			}
			if nums[inc] > nums[inc+1] {
				nums[inc], nums[inc+1] = nums[inc+1], nums[inc]
			}

			inc++
		}
		if inc+2 >= len(nums) {
			inc = 0
		}

	}
	return nums[0]

}