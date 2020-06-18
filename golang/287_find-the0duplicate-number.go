func findDuplicate(nums []int) int {
	var l, h, m, n int
	l = 1
	h = len(nums) - 1

	for l < h {
		m = (h + l) / 2
		n = 0
		for _, v := range nums {
			if v >= l && v <= m {
				n++
			}

		}
		if n > m-l+1 {
			h = m
		} else {
			l = m + 1
		}
	}
	return l

}