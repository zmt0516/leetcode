import "sort"

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)

	for _, v := range words {
		m[v]++
	}

	timesm := make(map[int][]string)

	for i, v := range m {
		timesm[v] = append(timesm[v], i)
		//times = append(times,v)
	}
	var times []int
	for i := range timesm {
		times = append(times, i)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(times)))
	var r []string
	for _, v := range times {
		sort.Strings(timesm[v])
		for _, v2 := range timesm[v] {
			if k > 0 {
				r = append(r, v2)
			} else {
				return r
			}
			k--
		}
	}

	return r

}