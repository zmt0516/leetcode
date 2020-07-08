func getArea(grid [][]int, i int, j int) int {

	r := 0
	if i < len(grid) && i >= 0 && j < len(grid[i]) && j >= 0 && grid[i][j] == 1 {
		grid[i][j] = 0
		r++
		r += getArea(grid, i, j+1)
		r += getArea(grid, i, j-1)
		r += getArea(grid, i+1, j)
		r += getArea(grid, i-1, j)

	}

	return r

}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := range grid {
		for j := range grid[i] {
			r := getArea(grid, i, j)
			if r > max {
				max = r
			}

		}

	}
	return max

}