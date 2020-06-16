func hammingWeight(num uint32) int {
	var n int
	for num > 0 {
		if num%2 == 1 {
			n += 1
		}
		num /= 2

	}
	return n
}