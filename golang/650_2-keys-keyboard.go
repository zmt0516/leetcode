func minSteps(n int) int {
	var r int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			r += i

		}
	}
	return r

}