var max int

func canFinish(numCourses int, prerequisites [][]int) bool {

	max = numCourses
	r := make([][]int, numCourses)
	emap := make(map[int]bool)

	for _, v := range prerequisites {

		r[v[0]] = append(r[v[0]], v[1])
		emap[v[0]] = true

	}

	for i := range emap {

		if emap[i] && finish(r, emap, 0, i) {
			return false
		}

	}

	//fmt.Println(r)
	return true

}

func finish(r [][]int, emap map[int]bool, n int, x int) bool {
	if n >= max {
		return true
	}
	for _, v := range r[x] {

		if emap[v] && finish(r, emap, n+1, v) {
			return true
		}
		emap[v] = false
	}
	return false

}