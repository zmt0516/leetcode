func singleNumber(nums []int) []int {
	var inc, inc2, x, y int
	for _, v := range nums {
		inc ^= v
	}
	inc2 = inc & (-inc)
	for _, v := range nums {
		if v&inc2 == 0 {
			x ^= v
		} else {
			y ^= v
		}

	}
	var r []int
	r = append(r, x, y)
	return r

}