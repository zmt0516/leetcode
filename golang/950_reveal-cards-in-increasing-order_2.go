import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	length := len(deck)
	if length < 3 {
		return deck
	}
	r := make([]int, length)
	r2 := make([]int, length)
	for i := 0; i < length; i++ {
		r[i] = i
	}

	chg := false
	i := 0

	for len(r) > 0 {
		j := r[0]
		r = r[1:]
		if !chg {
			r2[j] = deck[i]
			i++
		} else {
			r = append(r, j)
		}
		chg = !chg

	}
	return r2

}