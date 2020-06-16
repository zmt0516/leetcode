func reverse(x int) int {
	var r int
	for x != 0 {
		r *= 10
		r += x % 10
		x /= 10
	}
	maxint := int(^uint32(0) >> 1)
	if r > maxint || r < -1-maxint {
		return 0
	}

	return r

}