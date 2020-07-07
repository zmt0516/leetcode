func partitionLabels(S string) []int {
	var r []int
	lS := len(S)

	l := 0
	h := lS - 1
	mid := l
	omid := l
	for l < lS {
		has := make(map[byte]bool)
		mid = l
		for l < lS {
			if !has[S[l]] {
				has[S[l]] = true
				for h = lS - 1; h >= mid; h-- {
					if S[h] == S[l] {
						mid = h + 1
						break
					}

				}

			}
			l++
			if l == mid {
				break
			}

		}

		r = append(r, mid-omid)
		omid = mid

	}
	return r

}