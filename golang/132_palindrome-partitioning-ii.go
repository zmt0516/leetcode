
func minCut(s string) int {
	lens := len(s)
	if lens == 0 {
		return 0
	}

	emap := make([][]bool, lens)
	for i := range emap {
		emap[i] = make([]bool, lens)

	}

	for i := lens - 1; i >= 0; i-- {
		for j := lens - 1; j > i; j-- {
			if s[i] == s[j] {
				if i+1 == j {
					emap[i][j] = true
				} else {
					emap[i][j] = emap[i+1][j-1]

				}

			}

		}
		emap[i][i] = true
	}

	emap2 := make([]int, lens+1)
	for i := range emap2 {
		emap2[i] = 10000000
	}
	emap2[0] = 0
	for i := 0; i < lens; i++ {

		if emap[0][i] && emap2[i+1] > emap2[0] {
			emap2[i+1] = emap2[0]
		}

		for j := 1; j <= i; j++ {
			if emap[j][i] {
				if emap2[i+1] > emap2[j]+1 {
					emap2[i+1] = emap2[j] + 1
				}
			}
		}
	}

	return emap2[lens]

}




