func nthSuperUglyNumber(n int, primes []int) int {

	var r []int
	var pi []int
	maxint32 := int(^uint32(0) >> 1)
	r = append(r, 1)
	for i := range primes {
		pi = append(pi, 0*i)
	}
	for len(r) < n {
		ri := r[len(r)-1]
		for i, v := range primes {

			if r[pi[i]]*v <= ri {
				pi[i]++
			}
		}
		newr := maxint32
		for i, v := range primes {
			if r[pi[i]]*v < newr {
				newr = r[pi[i]] * v
			}
		}
		r = append(r, newr)

	}
	return r[n-1]

}