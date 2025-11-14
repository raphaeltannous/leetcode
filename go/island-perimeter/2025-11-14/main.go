func islandPerimeter(grid [][]int) int {
	result := 0

	seenGrid := make([][]bool, len(grid))
	for y, row := range grid {
		seenGrid[y] = make([]bool, len(row))
	}

	var dfs func(y, x int)
	dfs = func(y, x int) {
		if seenGrid[y][x] {
			return
		}

		directions := [][2]int{
			{y + 1, x},
			{y, x + 1},
			{y - 1, x},
			{y, x - 1},
		}

		for _, dir := range directions {
			if dir[0] < 0 || dir[0] > len(grid)-1 || dir[1] < 0 || dir[1] > len(grid[0])-1 {
				result += 1
				continue
			}

			if grid[dir[0]][dir[1]] == 0 {
				result += 1
			} else {
				seenGrid[y][x] = true
				dfs(dir[0], dir[1])
			}
		}
	}

	for y, row := range grid {
		for x, value := range row {
			if value == 1 {
				dfs(y, x)
				break
			}
		}
	}

	return result
}
