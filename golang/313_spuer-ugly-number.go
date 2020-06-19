import "sort"

func nthSuperUglyNumber(n int, primes []int) int {

	if n < 10000 {
		i := 1
		for ; n > 0; i++ {
			j := i
			for _, v := range primes {
				for j%v == 0 {
					j /= v

				}
				if j == 1 {
					n--
					break
				}
			}

		}
		return i - 1

	}

	var r []int
	maxint32 := int(^uint32(0) >> 1)
	r = append(r, 1)
	for _, v := range primes {

		for _, v2 := range r {
			vi := v
			for ; v2*vi < maxint32; vi *= v {
				r = append(r, v2*vi)

			}

		}

	}
	sort.Ints(r)
	return r[n-1]

}