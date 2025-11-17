func numIslands(grid [][]byte) int {
	seenGrid := make([][]bool, len(grid))
	for x, row := range grid {
		seenGrid[x] = make([]bool, len(row))
	}

	var dfs func(y, x int)
	dfs = func(y, x int) {
		if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 {
			return
		}

		if seenGrid[y][x] {
			return
		}

		if grid[y][x] == '0' {
			return
		} else {
			seenGrid[y][x] = true
		}

		dirs := [][2]int{
			{y + 1, x},
			{y - 1, x},
			{y, x + 1},
			{y, x - 1},
		}

		for _, dir := range dirs {
			dfs(dir[0], dir[1])
		}
	}

	islands := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == '1' && !seenGrid[y][x] {
				islands++
				dfs(y, x)
			}
		}
	}

	return islands
}
