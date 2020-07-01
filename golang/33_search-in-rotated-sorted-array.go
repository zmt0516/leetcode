
func findMinOrder(nums []int) int {
	left := 0
	right := len(nums) - 1
	if nums[right] > nums[left] {
		return left
	}

	for left < right-1 {
		mid := left + (right-left)/2
		if nums[mid] < nums[left] {
			right = mid

		} else {
			left = mid
		}
	}
	if left == right {
		return left
	}
	if nums[left] < nums[left+1] {
		return left
	}
	return left + 1

}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	min := findMinOrder(nums)
	nums = append(nums[min:], nums[:min]...)
	left := 0
	right := len(nums) - 1
	for left < right {
		//fmt.Println(left,right,min,nums)
		mid := left + (right-left)/2
		if nums[mid] == target {
			min += mid
			if min >= len(nums) {
				min -= len(nums)
			}
			return min
		}
		if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid - 1
		}
	}
	if nums[left] == target {
		min += left
		if min >= len(nums) {
			min -= len(nums)
		}
		return min
	}
	return -1

}