func minimumTotal(triangle [][]int) int {
	var r int
	if len(triangle) == 0 {
		return r
	}

	for i := 1; i < len(triangle); i++ {

		triangle[i][0] += triangle[i-1][0]

		for j := 1; j < i; j++ {
			if triangle[i-1][j-1] > triangle[i-1][j] {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += triangle[i-1][j-1]
			}

		}

		triangle[i][i] += triangle[i-1][i-1]

	}
	last := triangle[len(triangle)-1]
	r = last[0]
	for i := 1; i < len(triangle); i++ {
		if last[i] < r {
			r = last[i]
		}
	}
	return r

}