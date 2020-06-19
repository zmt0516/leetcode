import "sort"

func nthUglyNumber(n int) int {
	if n < 100 {
		for i := 1; ; i++ {
			j := i
			for j%2 == 0 {
				j /= 2
			}
			for j%3 == 0 {
				j /= 3
			}
			for j%5 == 0 {
				j /= 5
			}
			if j == 1 {
				n--
			}
			if n == 0 {
				return i
			}
		}
	}

	var r []int
	r2 := 1
	maxint32 := int(^uint32(0) >> 1)
	for i := 0; i < 32; i++ {
		r3 := 1
		for j := 0; j < 21; j++ {

			if r2*r3 > maxint32 {
				break
			}

			r5 := 1
			for k := 0; k < 14; k++ {

				if r2*r3*r5 > maxint32 {
					break
				}
				r = append(r, r2*r3*r5)

				r5 *= 5

			}
			r3 *= 3
		}
		r2 *= 2
	}
	sort.Ints(r)
	return r[n-1]
}