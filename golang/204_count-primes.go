func countPrimes(n int) int {
	var r []bool
	nr := 0
	for i := 0; i < n; i++ {
		r = append(r, true)
	}
	for i := 2; i < n; i++ {
		if r[i] {
			nr++
			for j := i + i; j < n; j += i {
				r[j] = false
			}

		}
	}

	return nr

}