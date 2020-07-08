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
	for i := range emap {
		if emap[i] == 0 {
			r := longest(envelopes, emap[:], i)
			if r > max {
				max = r
			}
		}
	}
	return max
}

func longest(envelopes [][]int, emap []int, i int) int {
	if emap[i] != 0 {
		return emap[i]
	}
	r := 0
	for j := i + 1; j < len(envelopes); j++ {
		if envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
			r2 := longest(envelopes, emap[:], j)
			if r2 > r {
				r = r2
			}
		}

	}
	return r + 1

}