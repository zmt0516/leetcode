
func firstMissingPositive(nums []int) int {

	n := 1
	n2 := 1
	inc := 0
	for len(nums) > 0 {
		if nums[inc] < n || nums[inc] > len(nums)+n-1 {
			nums = append(nums[:inc], nums[inc+1:]...)
		} else if nums[inc] == n {
			nums = append(nums[:inc], nums[inc+1:]...)
			n++
		} else {
			inc++
		}
		if inc > len(nums)-1 {
			if n2 == n {
				return n2
			}
			n2 = n
			inc = 0
		}

	}
	return n

}