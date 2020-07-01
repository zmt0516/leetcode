func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	mid := (len(nums) - 1) / 2
	if nums[mid] == target || nums[0] == target || nums[len(nums)-1] == target {
		return true
	}
	if len(nums) <= 3 {
		return false
	}

	if nums[0] > nums[mid] {
		if target > nums[0] || target < nums[mid] {

			return search(nums[1:mid], target)
		}
		return search(nums[mid+1:len(nums)-1], target)
	}
	if nums[0] < nums[mid] {
		//fmt.Println(nums[1:mid])
		if target > nums[0] && target < nums[mid] {
			return search(nums[1:mid], target)

		}
		return search(nums[mid+1:len(nums)-1], target)

	}
	if nums[0] == nums[mid] {
		return search(nums[mid+1:len(nums)-1], target) || search(nums[1:mid], target)
	}
	return false

}