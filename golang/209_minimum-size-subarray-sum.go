func minSubArrayLen(s int, nums []int) int {
	var r int
	lenn := len(nums)
	if s == 0 || lenn == 0 {
		return r
	}
	l := 0
	h := 0
	sum := 0
	for h = l; h < lenn; h++ {
		sum += nums[h]
		if sum >= s {
			r = h - l + 1
			break
		}

	}

	for {
		sum -= nums[l]
		l++
		if sum >= s {
			r--
		} else {

			h++
			if h >= lenn {
				break
			}
			sum += nums[h]
		}

	}
	return r

}