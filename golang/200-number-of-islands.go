func delislands(grid [][]byte, i int, j int) {

	if i < len(grid) && i >= 0 && j < len(grid[i]) && j >= 0 && grid[i][j] == '1' {
		grid[i][j] = '0'
		delislands(grid, i-1, j)
		delislands(grid, i+1, j)
		delislands(grid, i, j-1)
		delislands(grid, i, j+1)

	}

}

func numIslands(grid [][]byte) int {
	r := 0
	//fmt.Println(grid)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				r++
				delislands(grid, i, j)
			}

		}
	}

	return r

}