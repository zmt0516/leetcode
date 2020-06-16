func checkSubarraySum(nums []int, k int) bool {
	var yu []int
	var sum, h, l, m, lo int
	if k == 0 {
		for i, v := range nums {
			if i > 0 && nums[i-1]+v == 0 {
				return true

			}
		}
		return false
	}
	for i, v := range nums {
		sum += v
		if i == 0 {
			yu = append(yu, sum%k)
			continue
		}
		if sum%k == 0 {
			return true
		}
		lo = l
		h = i
		l = 0
		for l < h {
			m = (h + l) / 2
			if sum%k > yu[m] {
				l = m + 1
			} else {
				h = m
			}
		}
		if l < i && yu[l] == sum%k && (l != lo || l == lo && l+1 < i && yu[l+1] == sum%k) {
			return true
		}

		yu = append(yu, sum%k)
		copy(yu[l+1:], yu[l:])
		yu[l] = sum % k

	}
	return false

}