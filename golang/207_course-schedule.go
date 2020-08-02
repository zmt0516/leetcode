var max int

func canFinish(numCourses int, prerequisites [][]int) bool {

	max = numCourses
	r := make([][]int, numCourses)

	for _, v := range prerequisites {

		r[v[0]] = append(r[v[0]], v[1])

	}

	for i := range r {

		if finish(r, 0, i) {
			return false
		}

	}

	//fmt.Println(r)
	return true

}

func finish(r [][]int, n int, x int) bool {
	if n >= max {
		return true
	}
	for _, v := range r[x] {
		if finish(r, n+1, v) {
			return true
		}
	}
	return false

}