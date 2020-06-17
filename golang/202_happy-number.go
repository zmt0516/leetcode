func isHappy(n int) bool {
	r := n
	for i := 0; i < 100; i++ {
		for r = 0; n > 0; {
			r += (n % 10) * (n % 10)
			n /= 10

		}
		n = r
		if n == 1 {
			return true
		}

	}
	return false

}