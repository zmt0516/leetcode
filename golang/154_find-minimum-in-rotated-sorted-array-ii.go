func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1
	if nums[right] > nums[left] {
		return nums[left]
	}

	for left < right-1 {
		for left < right-1 && nums[left] == nums[left+1] {
			left++
		}
		for left < right-1 && nums[right] == nums[right-1] {
			right--
		}
		mid := left + (right-left)/2
		if nums[mid] < nums[left] {
			right = mid

		} else {
			left = mid
		}
	}
	if left == right {
		return nums[left]
	}
	if nums[left] < nums[left+1] {
		return nums[left]
	}
	return nums[left+1]

}