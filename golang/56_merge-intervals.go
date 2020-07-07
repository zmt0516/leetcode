import "sort"

type stype [][]int

func (a stype) Len() int { return len(a) }
func (a stype) Less(i, j int) bool {
	if a[i][0] < a[j][0] {
		return true
	}
	return false
}
func (a stype) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func merge(intervals [][]int) [][]int {
	var i int
	var k int

	sort.Sort(stype(intervals))
	for i < len(intervals) {
		v := intervals[i]
		for j := i + 1; j < len(intervals); {
			//fmt.Println(i,j,k)
			if intervals[j][0] <= v[1] {
				if v[1] < intervals[j][1] {
					v[1] = intervals[j][1]
				}

				i++
				j++

			} else {
				break

			}

		}
		//fmt.Println(i,k)
		intervals[k][0] = v[0]
		intervals[k][1] = v[1]
		i++
		k++

	}

	return intervals[:k]

}