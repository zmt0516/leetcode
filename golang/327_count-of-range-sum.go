func countRangeSum(nums []int, lower int, upper int) int {
	var n, sum, l, h, m int
	sums := []int{}
	for i, v := range nums {
		sum += v
		if sum >= lower && sum <= upper {
			n += 1
		}
		h = i
		l = 0
		for l < h {
			m = (h + l) / 2
			if sum-lower < sums[m] {
				h = m
			} else {
				l = m + 1
			}
		}
		n += l
		h = i
		l = 0
		for l < h {
			m = (h + l) / 2
			if sum-upper > sums[m] {
				l = m + 1
			} else {
				h = m
			}
		}
		n -= l
		h = i
		l = 0
		for l < h {
			m = (h + l) / 2
			if sum < sums[m] {
				h = m
			} else {
				l = m + 1
			}
		}
		sums = append(sums, sum)
		copy(sums[l+1:], sums[l:])
		sums[l] = sum

	}
	return n

}