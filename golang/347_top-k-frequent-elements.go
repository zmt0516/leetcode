import "sort"

func topKFrequent(nums []int, k int) []int {
	var r []int
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	keys := make(map[int][]int)
	for i, v := range m {
		keys[v] = append(keys[v], i)

	}
	var keys2 []int
	for i := range keys {
		keys2 = append(keys2, i)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys2)))
	i := 0
	for {
		for _, v := range keys[keys2[i]] {
			if k > 0 {
				r = append(r, v)
				k--

			}
			if k == 0 {
				return r
			}

		}
		i++

	}

	return r

}