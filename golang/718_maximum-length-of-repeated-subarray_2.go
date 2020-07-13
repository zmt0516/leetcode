func findLength(A []int, B []int) int {

	//lenA:=len(A)
	lenB := len(B)
	var r int

	emap := make([]int, lenB+1)

	for _, v := range A {

		for i := lenB; i > 0; i-- {

			if v == B[i-1] {
				lene := emap[i-1] + 1
				emap[i] = lene
				if lene > r {
					r = lene
				}
			} else {
				emap[i] = 0
			}

		}

	}
	return r

}
