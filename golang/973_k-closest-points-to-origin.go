import "sort"

func kClosest(points [][]int, K int) [][]int {

	m := make(map[int][][]int)

	for _, v := range points {
		m[v[0]*v[0]+v[1]*v[1]] = append(m[v[0]*v[0]+v[1]*v[1]], v)

	}
	var keys []int
	for i := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	var r [][]int
	for K > 0 {
		for _, v := range keys {
			for _, v2 := range m[v] {
				r = append(r, v2)
				K--
				if K == 0 {
					return r
				}
			}
		}
	}
	return r

}