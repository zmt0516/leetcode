func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	MaxUint := ^uint(0)
	h := int(MaxUint >> 1)
	l := h
	//l :=- h -1
	//r:= -MaxInt - 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < h && nums[i] > l {
			h = nums[i]
		}

		if nums[i] > nums[i-1] && nums[i] < h {
			l = nums[i-1]
			h = nums[i]
		}
		if nums[i] > h {
			return true
		}

	}
	return false

}