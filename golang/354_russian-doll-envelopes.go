import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i int, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		} else if envelopes[i][1] < envelopes[j][1] {
			return true
		}
		return false
	})
	emap := make([]int, len(envelopes))
	max := 0

	for i := range envelopes {
		r := 0
		for j := i - 1; j >= 0; j-- {
			if emap[j] > r && envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				r = emap[j]

			}

		}
		emap[i] = r + 1
		if r+1 > max {
			max = r + 1
		}

	}

	return max
}
