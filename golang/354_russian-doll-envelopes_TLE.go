// A stupid method which return a Time Limit Exceeded error.package golang.
// However, the results are true. Too Slow.
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
	MaxUint := ^uint(0)
	y := int(MaxUint >> 1)

	for i := range emap {
		if envelopes[i][1] < y && emap[i] == 0 {
			r := longest(envelopes, emap[:], i)
			if r >= max {
				y = envelopes[i][1]
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
	MaxUint := ^uint(0)
	y := int(MaxUint >> 1)
	for j := i + 1; j < len(envelopes); j++ {
		if envelopes[j][1] < y && envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
			r2 := longest(envelopes, emap[:], j)
			if r2 >= r {
				y = envelopes[j][1]
				r = r2
			}
		}

	}
	return r + 1

}