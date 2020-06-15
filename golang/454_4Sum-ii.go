import "sort"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)
	sort.Ints(D)
	N := len(A)
	n := 0
	var a, b, c, d int
	//    fh :=1

	for a = 0; a < N; a++ {
		for b = 0; b < N && A[a]+B[b]+C[0]+D[0] <= 0; b++ {
			if A[a]+B[b]+C[N-1]+D[N-1] >= 0 {
				for c = 0; c < N && A[a]+B[b]+C[c]+D[0] <= 0; c++ {
					if A[a]+B[b]+C[c]+D[N-1] >= 0 {
						if A[a]+B[b]+C[c]+D[d] <= 0 {
							if A[a]+B[b]+C[c]+D[d] == 0 {
								for ; d >= 0; d-- {
									if A[a]+B[b]+C[c]+D[d] < 0 {
										break
									}

								}
								if d < 0 {
									d += 1
								}

							}
							for ; d < N; d++ {
								if A[a]+B[b]+C[c]+D[d] == 0 {
									n += 1

								}
								if A[a]+B[b]+C[c]+D[d] > 0 {
									break
								}

							}
							if d == N {
								d -= 1
							}

						} else {
							for ; d >= 0; d-- {
								if A[a]+B[b]+C[c]+D[d] == 0 {
									n += 1
								}
								if A[a]+B[b]+C[c]+D[d] < 0 {
									break
								}
							}
							if d == -1 {
								d += 1
							}

						}

					}

				}

			}

		}

	}

	return n

}