func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	max := 1
	emap := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {

		r := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= nums[i] && emap[j] == r+1 {
				break
			}
			if r < emap[j] && nums[j] < nums[i] {
				r = emap[j]
			}
		}
		emap[i] = r + 1
		if r+1 > max {
			max = r + 1
		}

	}

	return max

}