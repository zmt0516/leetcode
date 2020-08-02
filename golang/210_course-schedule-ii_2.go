func findOrder(numCourses int, prerequisites [][]int) []int {

	emap := make([][]int, numCourses)
	emap2 := make([]int, numCourses)

	for _, v := range prerequisites {

		emap[v[1]] = append(emap[v[1]], v[0])
		emap2[v[0]]++

	}
	var r []int
	var r2 []int
	for i, v := range emap2 {
		if v == 0 {
			r2 = append(r2, i)
		}

	}

	for len(r2) > 0 {
		v := r2[0]
		r2 = r2[1:]
		r = append(r, v)
		for _, v2 := range emap[v] {
			emap2[v2]--
			if emap2[v2] == 0 {
				r2 = append(r2, v2)
			}
		}

	}
	if len(r) != numCourses {
		r = make([]int, 0)
	}
	return r

}