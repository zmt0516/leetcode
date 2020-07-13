func findLength(A []int, B []int) int {

	lenA := len(A)
	lenB := len(B)
	var r int

	for i := 0; i < lenA && i+r < lenA; i++ {
	LABEL:
		for j := 0; j < lenB && j+r < lenB; j++ {
			k := 0
			for ; i+k < lenA && j+k < lenB; k++ {
				if A[i+k] != B[j+k] {
					if k > r {
						r = k
					}
					continue LABEL
				}

			}
			if k > r {
				r = k
			}

		}

	}

	return r

}
