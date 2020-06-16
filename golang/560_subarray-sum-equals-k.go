func subarraySum(nums []int, k int) int {

	var n, sum, l, h, m int
	sums := []int{}
	for i, v := range nums {
		sum += v
		if sum == k {
			n += 1
		}

		h = i
		l = 0
		for l < h {
			m = (h + l) / 2
			if sum-k > sums[m] {
				l = m + 1
			} else {
				h = m
			}
		}
		for ; l < i; l++ {
			if sums[l] == sum-k {
				n += 1
			} else {
				break
			}
		}

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