import "sort"

func findOrder(numCourses int, prerequisites [][]int) []int {

	var r []int

	if !canFinish(numCourses, prerequisites) {
		return r
	}

	var n int

	emap := make([]int, numCourses)
	//n1=numCourses+1
	for {

		for _, v := range prerequisites {

			if emap[v[0]] <= emap[v[1]] {
				emap[v[1]] = emap[v[0]] - 1
				n++
			}

		}
		if n == 0 {
			//fmt.Println("OK!")
			break
		}
		n = 0
	}
	var emap2 [][]int
	for i, v := range emap {
		emap2 = append(emap2, []int{v, i})

	}
	//fmt.Println(emap2)
	sort.Slice(emap2, func(p, q int) bool {
		return emap2[p][0] < emap2[q][0]
	})
	for _, v := range emap2 {
		r = append(r, v[1])
	}

	return r

}

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


