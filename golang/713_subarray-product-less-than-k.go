func numSubarrayProductLessThanK(nums []int, k int) int {

	n := 0
	length := len(nums)
	p := 1

	h := 0
	l := 0
	for h < length {
		p *= nums[h]
		if p < k {
			n += h - l + 1
			h++

		} else {
			for p >= k && l <= h {
				p /= nums[l]
				l++
			}
			n += h - l + 1
			h++

		}

	}

	return n

}
