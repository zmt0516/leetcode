func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	if newInterval[0] > intervals[len(intervals)-1][1] {

		intervals = append(intervals, newInterval)
		return intervals
	}

	var r [][]int
	var i int
	for i = range intervals {
		v := intervals[i]
		if newInterval[0] > v[1] {

			r = append(r, v)

		} else {
			break

		}

	}
	if intervals[i][0] < newInterval[0] {
		newInterval[0] = intervals[i][0]
	}
	//fmt.Println(i)
	for ; i < len(intervals); i++ {
		v := intervals[i]
		//fmt.Println(i,"!")
		if v[0] <= newInterval[1] {
			if v[1] > newInterval[1] {
				newInterval[1] = v[1]

			}

		} else {

			break
		}

	}
	//fmt.Println(i)
	r = append(r, newInterval)

	for ; i < len(intervals); i++ {
		v := intervals[i]
		r = append(r, v)

	}
	return r

}