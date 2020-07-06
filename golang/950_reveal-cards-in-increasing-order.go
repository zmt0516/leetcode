import "sort"

func deckRevealedIncreasing(deck []int) []int {
	var r []int
	sort.Ints(deck)
	for len(deck) != 0 {
		pop := deck[len(deck)-1]
		deck = deck[:len(deck)-1]

		if len(r) != 0 {

			r = append([]int{pop, r[len(r)-1]}, r[:len(r)-1]...)

		} else {
			r = append(r, pop)

		}

	}
	return r

}