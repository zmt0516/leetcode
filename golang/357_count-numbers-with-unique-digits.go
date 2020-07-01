func countNumbersWithUniqueDigits(n int) int {
	r := 1
	rn := 9
	for i := 1; i <= 10 && i <= n; i++ {

		r += rn
		rn *= 10 - i

	}
	return r

}